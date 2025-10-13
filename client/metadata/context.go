package metadata

import (
	"context"

	gogrpc "github.com/ralvarezdev/go-grpc"
	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

// GetCtxMetadata metadata from the context
//
// Parameters:
//
//   - ctx: The context to get the metadata from
//
// Returns:
//
//   - metadata.MD: The metadata from the context
//   - error: An error if the metadata is not found
func GetCtxMetadata(ctx context.Context) (metadata.MD, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return nil, status.Error(
			codes.Aborted,
			gojwtgrpc.ErrMissingMetadata.Error(),
		)
	}
	return md, nil
}

// GetCtxMetadataBearerToken gets the bearer token from the metadata
//
// Parameters:
//
//   - ctx: The context to get the metadata from
//   - key: The metadata key where the token is stored
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetCtxMetadataBearerToken(ctx context.Context, key string) (
	string,
	error,
) {
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return gogrpcmd.GetMetadataBearerToken(md, key)
}

// GetCtxMetadataAuthorizationToken gets the authorization token from the context metadata
//
// Parameters:
//
//   - ctx: The context to get the metadata from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetCtxMetadataAuthorizationToken(ctx context.Context) (string, error) {
	return GetCtxMetadataBearerToken(ctx, gogrpc.AuthorizationMetadataKey)
}

// GetCtxMetadataGCloudAuthorizationToken gets the GCloud authorization token from the context metadata
//
// Parameters:
//
//   - ctx: The context to get the metadata from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetCtxMetadataGCloudAuthorizationToken(ctx context.Context) (
	string,
	error,
) {
	return GetCtxMetadataBearerToken(ctx, gogrpc.GCloudAuthorizationMetadataKey)
}
