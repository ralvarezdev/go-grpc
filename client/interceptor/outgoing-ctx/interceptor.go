package outgoing_ctx

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type (
	// Interceptor is the interceptor for the outgoing context
	Interceptor struct {
		logger *slog.Logger
	}
)

// NewInterceptor creates a new interceptor for the outgoing context
//
// Parameters:
//
//   - logger: the logger to use
//
// Returns:
//
//   - *Interceptor: the interceptor
func NewInterceptor(logger *slog.Logger) *Interceptor {
	if logger != nil {
		logger = logger.With(
			slog.String("component", "client_interceptor_outgoing_ctx"),
		)
	}

	return &Interceptor{
		logger,
	}
}

// PrintOutgoingCtx prints the outgoing context
//
// Returns:
//
//   - grpc.UnaryClientInterceptor: the interceptor
func (i Interceptor) PrintOutgoingCtx() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Get the outgoing context
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			return status.Error(
				codes.Internal,
				ErrFailedToGetOutgoingContext.Error(),
			)
		}

		// Print the metadata
		if i.logger != nil {
			for key, values := range md {
				for _, value := range values {
					i.logger.Debug(
						"Found metadata in outgoing context",
						slog.String("method", method),
						slog.String("key", key),
						slog.String("value", value),
					)
				}
			}
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
