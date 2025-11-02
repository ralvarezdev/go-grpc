package jwt

import (
	"context"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
)

type (
	// Interceptor is the interceptor for the authentication
	Interceptor struct {
		gCloudAccessToken string
		logger            *slog.Logger
	}
)

// NewInterceptor creates a new authentication interceptor
//
// Parameters:
//
//   - gCloudTokenSource: the GCloud token source to get the access token from
//   - logger: the logger to use for logging
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: an error if the token source or any other error occurs
func NewInterceptor(
	gCloudTokenSource *oauth.TokenSource,
	logger *slog.Logger,
) (*Interceptor, error) {
	// Get the access token from the token source
	token, err := gCloudTokenSource.Token()
	if err != nil {
		return nil, err
	}

	if logger != nil {
		logger = logger.With(
			slog.String(
				"grpc_client_interceptor",
				"gcloud_authenticator",
			),
		)
	}

	return &Interceptor{
		gCloudAccessToken: token.AccessToken,
		logger:            logger,
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
		// Add GCloud authorization token to the outgoing context metadata
		ctx, err := gogrpcmd.SetOutgoingCtxMetadataGCloudAuthorizationToken(
			ctx,
			i.gCloudAccessToken,
		)
		if err != nil {
			if i.logger != nil {
				i.logger.Warn(
					"Failed to set GCloud metadata authorization token for the gRPC client",
					slog.String("error", err.Error()),
				)
			}
		}

		// Invoke the original invoker
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
