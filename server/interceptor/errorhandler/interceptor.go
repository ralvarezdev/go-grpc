package errorhandler

import (
	"context"
	"log/slog"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	goflagsmode	"github.com/ralvarezdev/go-flags/mode"
	goflags "github.com/ralvarezdev/go-flags"

	gogrpc "github.com/ralvarezdev/go-grpc"
)

type (
	// Interceptor is the interceptor for the error handler
	Interceptor struct {
		modeFlag *goflagsmode.Flag
		logger *slog.Logger
	}
)

// NewInterceptor creates a new error handler interceptor
//
// Parameters:
//
//   - modeFlag: the application mode flag
//   - logger: the logger to use (can be nil)
//
// Returns:
//
//  - *Interceptor: the interceptor
//  - error: if there was an error creating the interceptor
func NewInterceptor(modeFlag *goflagsmode.Flag, logger *slog.Logger) (*Interceptor, error) {
	// Check if the mode flag is nil
	if modeFlag == nil {
		return nil, goflags.ErrNilFlag
	}
	
	// Create the logger for the interceptor
	if logger != nil {
		logger = logger.With(
			slog.String("grpc_client_interceptor", "error_handler"),
		)
	}

	return &Interceptor{
		modeFlag: modeFlag,
		logger: logger,
	}, nil
}

// HandleError returns the error handler interceptor
//
// Returns:
//
//   - grpc.UnaryServerInterceptor: the error handler interceptor
func (i Interceptor) HandleError() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (value any, err error) {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic
				stack := debug.Stack()
				if i.logger != nil {
					i.logger.Error(
						"Panic recovered",
						slog.Any("method", info.FullMethod),
						slog.Any("error", r),
						slog.String("stack_trace", string(stack)),
					)
				}

				// Check if we are in production mode
				if i.modeFlag.IsProd() {
					// Set the error to internal server error
					err = status.Error(codes.Internal, gogrpc.InternalServerError)
				} else {
					// Set the error to the panic message
					err = status.Errorf(
						codes.Internal,
						"Panic: %v\nStack Trace:\n%s",
						r,
						string(stack),
					)
				}
			}
		}()
		return handler(ctx, req)
	}
}
