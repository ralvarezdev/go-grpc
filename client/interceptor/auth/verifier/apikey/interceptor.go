package apikey

import (
	"context"
	"log/slog"

	goapikeygrpc "github.com/ralvarezdev/go-api-key/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
)

type (
	// Interceptor is the interceptor for the authentication
	Interceptor struct {
		interceptions map[string]struct{}
		logger        *slog.Logger
	}
)

// NewInterceptor creates a new authentication interceptor
//
// Parameters:
//
//   - methodsToIntercept: a slice of method names to intercept
//   - logger: the logger to use for logging
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: an error if the interceptions map is nil
func NewInterceptor(
	methodsToIntercept []string,
	logger *slog.Logger,
) (*Interceptor, error) {
	// Check if the gRPC interceptions is nil
	if methodsToIntercept == nil {
		return nil, goapikeygrpc.ErrNilGRPCInterceptions
	}

	// Create a map of methods to intercept for efficient lookup
	interceptions := make(map[string]struct{})
	for _, method := range methodsToIntercept {
		interceptions[method] = struct{}{}
	}

	if logger != nil {
		logger = logger.With(
			slog.String(
				"grpc_client_interceptor",
				"api_key_verifier",
			),
		)
	}

	return &Interceptor{
		interceptions: interceptions,
		logger:        logger,
	}, nil
}

// Verify returns a new unary client interceptor that verifies the authentication metadata from the context is set if
// needed
//
// Returns:
//
//   - grpc.UnaryClientInterceptor: the interceptor
func (i Interceptor) Verify() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Check if the method should be intercepted
		_, ok := i.interceptions[method]

		// If the method is intercepted, verify it has the authorization metadata
		if ok {
			// Try to get the authorization metadata from the context
			_, err := gogrpcmd.GetCtxMetadataAuthorizationToken(
				ctx,
			)
			if err != nil {
				if i.logger != nil {
					i.logger.Warn(
						"Missing authorization metadata for intercepted method",
						slog.String("method", method),
						slog.String("error", err.Error()),
					)
				}
				return status.Errorf(
					codes.Unauthenticated,
					"Missing authorization metadata for intercepted method: %s",
					method,
				)
			}
		}

		// Invoke the original invoker
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
