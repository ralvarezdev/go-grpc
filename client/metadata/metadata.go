package metadata

import (
	"context"

	gogrpc "github.com/ralvarezdev/go-grpc"
	gojwt "github.com/ralvarezdev/go-jwt"
	"google.golang.org/grpc/metadata"
)

// SetCtxBearerToken sets the bearer token in the context metadata
//
// Parameters:
//
//   - ctx: The context to set the bearer token in
//   - key: The key to set the bearer token for
//   - token: The token to set in the context
//
// Returns:
//
//   - context.Context: The context with the bearer token set
func SetCtxBearerToken(ctx context.Context, key, token string) context.Context {
	return metadata.AppendToOutgoingContext(
		ctx,
		key,
		gojwt.BearerPrefix+" "+token,
	)
}

// SetCtxAuthorization sets the authorization in the context metadata
//
// Parameters:
//
//   - ctx: The context to set the authorization in
//   - token: The token to set in the context
//
// Returns:
//
//   - context.Context: The context with the authorization set
func SetCtxAuthorization(ctx context.Context, token string) context.Context {
	return SetCtxBearerToken(ctx, gogrpc.AuthorizationMetadataKey, token)
}

// SetCtxGCloudAuthorization sets the GCloud authorization in the context metadata
//
// Parameters:
//
//   - ctx: The context to set the GCloud authorization in
//   - gcloudToken: The GCloud token to set in the context
//
// Returns:
//
//   - context.Context: The context with the GCloud authorization set
func SetCtxGCloudAuthorization(
	ctx context.Context,
	gcloudToken string,
) context.Context {
	return SetCtxBearerToken(
		ctx,
		gogrpc.GCloudAuthorizationMetadataKey,
		gcloudToken,
	)
}
