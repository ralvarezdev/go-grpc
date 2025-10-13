package auth

import (
	"context"

	gogrpcclientmd "github.com/ralvarezdev/go-grpc/client/metadata"
	gogrpcservermd "github.com/ralvarezdev/go-grpc/metadata"
	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type (
	// Interceptor is the interceptor for the authentication
	Interceptor struct {
		interceptions     map[string]*gojwttoken.Token
		GCloudAccessToken *string
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
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: an error if the token source or the gRPC interceptions is nil or any other error occurs
func NewInterceptor(
	interceptions map[string]*gojwttoken.Token,
	options *Options,
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

	return &Interceptor{
		GCloudAccessToken: gCloudAccessToken,
		interceptions:     interceptions,
	}, nil
}

// Authenticate returns a new unary client interceptor that adds authentication metadata to the context
//
// Returns:
//
//   - grpc.UnaryClientInterceptor: the interceptor
func (i Interceptor) Authenticate() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Check if the method should be intercepted
		interception, ok := i.interceptions[method]
		if !ok || interception == nil {
			// Create the unauthenticated context metadata if the access token is not nil
			if i.GCloudAccessToken != nil {
				ctx = gogrpcclientmd.SetCtxGCloudAuthorization(
					ctx,
					*i.GCloudAccessToken,
				)
			}
		} else {
			// Get metadata from the context
			md, ok := metadata.FromOutgoingContext(ctx)
			if !ok {
				return status.Error(
					codes.Unauthenticated,
					gojwtgrpc.ErrMissingMetadata.Error(),
				)
			}

			// Get the raw token from the metadata
			rawToken, err := gogrpcservermd.GetMetadataAuthorizationToken(md)
			if err != nil {
				return status.Error(codes.Unauthenticated, err.Error())
			}

			// Create the authenticated context metadata
			ctx = gogrpcclientmd.SetCtxAuthorization(ctx, rawToken)
			if i.GCloudAccessToken == nil {
				ctx = gogrpcclientmd.SetCtxGCloudAuthorization(
					ctx,
					*i.GCloudAccessToken,
				)
			}
		}

		// Invoke the original invoker
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
