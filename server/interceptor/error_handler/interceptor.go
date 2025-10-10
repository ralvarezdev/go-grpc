package error_handler

import (
	"context"
	"log/slog"

	gogrpc "github.com/ralvarezdev/go-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	// Interceptor is the interceptor for the error handler
	Interceptor struct {
		logger *slog.Logger
	}
)

// NewInterceptor creates a new error handler interceptor
//
// Parameters:
//
//   - logger: the logger to use (can be nil)
//
// Returns:
//
//   - *Interceptor: the interceptor
func NewInterceptor(logger *slog.Logger) *Interceptor {
	if logger != nil {
		logger = logger.With(
			slog.String("interceptor", "error_handler"),
		)
	}

	return &Interceptor{
		logger: logger,
	}
}

// HandleError returns the error handler interceptor
//
// Returns:
//
//   - grpc.UnaryServerInterceptor: the error handler interceptor
func (i Interceptor) HandleError() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (value interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic
				if i.logger != nil {
					i.logger.Error(
						"panic recovered",
						slog.Any("method", info.FullMethod),
						slog.Any("error", r),
					)
				}

				// Set the error to internal server error
				err = status.Errorf(codes.Internal, gogrpc.InternalServerError)
			}
		}()
		return handler(ctx, req)
	}
}
