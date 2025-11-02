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
		logger        *slog.Logger
		apiKey        string
		interceptions map[string]struct{}
	}
)

// NewInterceptor creates a new authentication interceptor
//
// Parameters:
//
//   - methodsToIntercept: a slice of method names to intercept
//   - apiKey: the API key to use for authentication
//   - logger: the logger to use for logging
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: an error if the interceptions map is nil
func NewInterceptor(
	methodsToIntercept []string,
	apiKey string,
	logger *slog.Logger,
) (*Interceptor, error) {
	// Check if the gRPC interceptions is nil
	if methodsToIntercept == nil {
		return nil, goapikeygrpc.ErrNilGRPCInterceptions
	}

	// Check if the API key is empty
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
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
				"api_key_authenticator",
			),
		)
	}

	return &Interceptor{
		interceptions: interceptions,
		apiKey:        apiKey,
		logger:        logger,
	}, nil
}

// Authenticate creates a gRPC unary client interceptor for authentication
//
// Returns:
//
//   - grpc.UnaryClientInterceptor: the interceptor
func (i Interceptor) Authenticate() grpc.UnaryClientInterceptor {
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

		// Invoke the original invoker
		if !ok {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		// Set context metadata for the gRPC client with the API key
		ctx, err := gogrpcmd.SetOutgoingCtxMetadataAuthorizationToken(
			ctx,
			i.apiKey,
		)
		if err != nil {
			if i.logger != nil {
				i.logger.Error(
					"Failed to set metadata authorization token for the gRPC client",
					slog.String("error", err.Error()),
				)
			}
			return status.Error(
				codes.Internal,
				ErrFailedToSetMetadataAuthorizationToken.Error(),
			)
		}

		// Invoke the original invoker with the updated context
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
