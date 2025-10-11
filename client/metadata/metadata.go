package metadata

import (
	"context"

	gogrpcgcloud "github.com/ralvarezdev/go-grpc/cloud/gcloud"
	gogrpcservermd "github.com/ralvarezdev/go-grpc/server/metadata"
	gojwt "github.com/ralvarezdev/go-jwt"
	"google.golang.org/grpc/metadata"
)

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
	return metadata.AppendToOutgoingContext(
		ctx,
		gogrpcservermd.AuthorizationKey,
		gojwt.BearerPrefix+" "+token,
	)
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
	return metadata.AppendToOutgoingContext(
		ctx,
		gogrpcgcloud.AuthorizationMetadataKey,
		gojwt.BearerPrefix+" "+gcloudToken,
	)
}
