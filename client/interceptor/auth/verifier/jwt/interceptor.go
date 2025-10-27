package jwt

import (
	"context"
	"log/slog"

	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/status"

	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
)

type (
	// Interceptor is the interceptor for the authentication
	Interceptor struct {
		interceptions     map[string]*gojwttoken.Token
		gCloudAccessToken *string
		logger            *slog.Logger
	}

	// Options is the options for the interceptor
	Options struct {
		GCloudTokenSource *oauth.TokenSource
	}
)

// NewOptions creates a new options for the interceptor
//
// Parameters:
//
//   - gCloudTokenSource: the OAuth token source to get the access token for Google Cloud services
//
// Returns:
//
//   - *Options: the options
func NewOptions(
	gCloudTokenSource *oauth.TokenSource,
) *Options {
	return &Options{
		GCloudTokenSource: gCloudTokenSource,
	}
}

// NewInterceptor creates a new authentication interceptor
//
// Parameters:
//
//   - interceptions: the gRPC interceptions to determine which methods require authentication
//   - options: the options for the interceptor
//   - logger: the logger to use for logging
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: an error if the token source or the gRPC interceptions is nil or any other error occurs
func NewInterceptor(
	interceptions map[string]*gojwttoken.Token,
	options *Options,
	logger *slog.Logger,
) (*Interceptor, error) {
	// Check if the gRPC interceptions is nil
	if interceptions == nil {
		return nil, gojwtgrpc.ErrNilGRPCInterceptions
	}

	// Initialize the access token variable
	var gCloudAccessToken *string
	if options != nil && options.GCloudTokenSource != nil {
		// Get the access token from the token source
		token, err := options.GCloudTokenSource.Token()
		if err != nil {
			return nil, err
		}

		// Set the access token
		gCloudAccessToken = &token.AccessToken
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
		gCloudAccessToken: gCloudAccessToken,
		interceptions:     interceptions,
		logger:            logger,
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
		interception, ok := i.interceptions[method]

		// Add GCloud authorization if available
		var err error
		if i.gCloudAccessToken != nil {
			ctx, err = gogrpcmd.SetCtxMetadataGCloudAuthorizationToken(
				ctx,
				*i.gCloudAccessToken,
			)
			if err != nil {
				if i.logger != nil {
					i.logger.Warn(
						"Failed to set GCloud metadata authorization token for the gRPC client",
						slog.String("error", err.Error()),
					)
				}
			}
		}

		// If the method is intercepted, verify it has the authorization metadata
		if ok && interception != nil {
			// Try to get the authorization metadata from the context
			_, authErr := gogrpcmd.GetCtxMetadataAuthorizationToken(
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
