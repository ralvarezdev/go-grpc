package jwt

import (
	"context"
	"log/slog"

	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gogrpc "github.com/ralvarezdev/go-grpc"

	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
)

type (
	// Interceptor is the interceptor for the authentication
	Interceptor struct {
		interceptions map[string]*gojwttoken.Token
		logger        *slog.Logger
	}
)

// NewInterceptor creates a new authentication interceptor
//
// Parameters:
//
//   - interceptions: the gRPC interceptions to determine which methods require authentication
//   - options: the options for the interceptor
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: an error if the gRPC interceptions is nil or any other error occurs
func NewInterceptor(
	interceptions map[string]*gojwttoken.Token,
	logger *slog.Logger,
) (*Interceptor, error) {
	// Check if the gRPC interceptions is nil
	if interceptions == nil {
		return nil, gogrpc.ErrNilInterceptions
	}

	if logger != nil {
		logger = logger.With(
			slog.String(
				"grpc_client_interceptor",
				"jwt_verifier",
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
		// Check if the method should be intercepted, if so, verify the authorization metadata is set
		interception, ok := i.interceptions[method]
		if !ok {
			// Log the error and return an internal server error
			if i.logger != nil {
				i.logger.Error(
					"Could not find interception for method",
					slog.String("method", method),
				)
			}
			
			return status.Errorf(
				codes.Internal,
				gogrpc.InternalServerError,
				method,
			)
		}
		if interception == nil {
			// Try to get the authorization metadata from the context
			_, authErr := gogrpcmd.GetOutgoingCtxMetadataAuthorizationToken(
				ctx,
			)
			if authErr != nil {
				if i.logger != nil {
					i.logger.Warn(
						"Missing authorization metadata for intercepted method",
						slog.String("method", method),
						slog.String("interception", interception.String()),
						slog.String("error", authErr.Error()),
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
