package metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// GetIncomingCtxMetadata gets the incoming metadata from the context
//
// Parameters:
//
//   - ctx: The context to get the incoming metadata from
//
// Returns:
//
//   - metadata.MD: The incoming metadata from the context
//   - error: An error if the metadata is not found or any other error occurs
func GetIncomingCtxMetadata(ctx context.Context) (metadata.MD, error) {
	// Get the metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrNilMetadata
	}
	return md, nil
}

// GetOutgoingCtxMetadata gets the outgoing metadata from the context
//
// Parameters:
//
//   - ctx: The context to get the outgoing metadata from
//
// Returns:
//
// - metadata.MD: The outgoing metadata from the context
func GetOutgoingCtxMetadata(ctx context.Context) metadata.MD {
	// Get the metadata from the context
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	return md
}

// GetIncomingCtxMetadataValue gets the value for a given key from the context metadata
//
// Parameters:
//
//   - ctx: The context to get the metadata from
//   - key: The key to get the value for
//
// Returns:
//
//   - []string: The value for the given key
//   - error: An error if the key is not found or any other error occurs
func GetIncomingCtxMetadataValue(ctx context.Context, key string) ([]string, error) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	return GetMetadataValue(md, key)
}

// DeleteIncomingCtxMetadataValue deletes the value for a given key from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to get the metadata from
//   - key: The key to delete the value for
//
// Returns:
//
//   - context.Context: The context with the value deleted
//   - error: An error if the key is not found or any other error occurs
func DeleteIncomingCtxMetadataValue(ctx context.Context, key string) (
	context.Context,
	error,
) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = DeleteMetadataValue(md, key)
	return metadata.NewIncomingContext(ctx, md), nil
}

// GetIncomingCtxMetadataBearerToken gets the bearer token from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to get the metadata from
//   - key: The key to get the token for
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetIncomingCtxMetadataBearerToken(ctx context.Context, key string) (
	string,
	error,
) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataBearerToken(md, key)
}

// GetIncomingCtxMetadataAuthorizationToken gets the authorization token from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to get the metadata from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetIncomingCtxMetadataAuthorizationToken(ctx context.Context) (string, error) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataAuthorizationToken(md)
}

// GetIncomingCtxMetadataGCloudAuthorizationToken gets the GCloud authorization token from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to get the metadata from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetIncomingCtxMetadataGCloudAuthorizationToken(ctx context.Context) (
	string,
	error,
) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataGCloudAuthorizationToken(md)
}

// GetIncomingCtxMetadataRefreshToken gets the refresh token from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetIncomingCtxMetadataRefreshToken(ctx context.Context) (string, error) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataRefreshToken(md)
}

// GetOutgoingCtxMetadataAuthorizationToken gets the authorization token from the outgoing context metadata
//
// Parameters:
//
// - ctx: The outgoing context to get the token from
//
// Returns:
//
// - string: The token
// - error: An error if the token is not found or any other error occurs
func GetOutgoingCtxMetadataAuthorizationToken(ctx context.Context) (
	string,
	error,
) {
	// Get the metadata from the context
	md := GetOutgoingCtxMetadata(ctx)
	return GetMetadataAuthorizationToken(md)
}

// GetOutgoingCtxMetadataGCloudAuthorizationToken gets the GCloud authorization token from the outgoing context metadata
//
// Parameters:
//
// - ctx: The outgoing context to get the token from
//
// Returns:
//
// - string: The token
// - error: An error if the token is not found or any other error occurs
func GetOutgoingCtxMetadataGCloudAuthorizationToken(ctx context.Context) (
	string,
	error,
) {
	// Get the metadata from the context
	md := GetOutgoingCtxMetadata(ctx)
	return GetMetadataGCloudAuthorizationToken(md)
}

// GetOutgoingCtxMetadataRefreshToken gets the refresh token from the outgoing context metadata
//
// Parameters:
//
//   - ctx: The outgoing context to get the token from
//
// Returns:
//
// - string: The token
// - error: An error if the token is not found or any other error occurs
func GetOutgoingCtxMetadataRefreshToken(ctx context.Context) (string, error) {
	// Get the metadata from the context
	md := GetOutgoingCtxMetadata(ctx)
	return GetMetadataRefreshToken(md)
}

// GetOutgoingCtxMetadataAccessToken gets the access token from the outgoing context metadata
//
// Parameters:
//
// - ctx: The outgoing context to get the token from
//
// Returns:
//
// - string: The token
// - error: An error if the token is not found or any other error occurs
func GetOutgoingCtxMetadataAccessToken(ctx context.Context) (string, error) {
	// Get the metadata from the context
	md := GetOutgoingCtxMetadata(ctx)
	return GetMetadataAccessToken(md)
}

// SetOutgoingCtxMetadataBearerToken sets the authorization token to the metadata
//
// Parameters:
//
//   - ctx: The outgoing context to set the token to
//   - key: The metadata key where the token will be set
//   - token: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
//   - error: An error if the metadata is not found or any other error occurs
func SetOutgoingCtxMetadataBearerToken(
	ctx context.Context,
	key, token string,
) (context.Context, error) {
	md := SetMetadataBearerToken(GetOutgoingCtxMetadata(ctx), key, token)
	return metadata.NewIncomingContext(ctx, md), nil
}

// SetOutgoingCtxMetadataAuthorizationToken sets the authorization token to the metadata
//
// Parameters:
//
//   - ctx: The outgoing context to set the token to
//   - token: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
func SetOutgoingCtxMetadataAuthorizationToken(
	ctx context.Context,
	token string,
) (context.Context, error) {
	md := SetMetadataAuthorizationToken(GetOutgoingCtxMetadata(ctx), token)
	return metadata.NewIncomingContext(ctx, md), nil
}

// SetOutgoingCtxMetadataGCloudAuthorizationToken sets the GCloud authorization token to the metadata
//
// Parameters:
//
//   - ctx: The outgoing context to set the token to
//   - token: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
//   - error: An error if the metadata is not found or any other error occurs
func SetOutgoingCtxMetadataGCloudAuthorizationToken(
	ctx context.Context,
	token string,
) (context.Context, error) {
	md := SetMetadataGCloudAuthorizationToken(
		GetOutgoingCtxMetadata(ctx),
		token,
	)
	return metadata.NewIncomingContext(ctx, md), nil
}

// SetOutgoingCtxMetadataRefreshToken sets the refresh token to the metadata
//
// Parameters:
//
//   - ctx: The outgoing context to set the token to
//   - refreshToken: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
//   - error: An error if the metadata is not found or any other error occurs
func SetOutgoingCtxMetadataRefreshToken(
	ctx context.Context,
	refreshToken string,
) (context.Context, error) {
	md := SetMetadataRefreshToken(
		GetOutgoingCtxMetadata(ctx),
		refreshToken,
	)
	return metadata.NewIncomingContext(ctx, md), nil
}

// SetOutgoingCtxMetadataAccessToken sets the access token to the metadata
//
// Parameters:
//
//   - ctx: The outgoing context to set the token to
//   - accessToken: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
//   - error: An error if the metadata is not found or any other error occurs
func SetOutgoingCtxMetadataAccessToken(
	ctx context.Context,
	accessToken string,
) (context.Context, error) {
	md := SetMetadataAccessToken(
		GetOutgoingCtxMetadata(ctx),
		accessToken,
	)
	return metadata.NewIncomingContext(ctx, md), nil
}

// GetIncomingCtxMetadataAccessToken gets the access token from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetIncomingCtxMetadataAccessToken(ctx context.Context) (string, error) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataAccessToken(md)
}

// ClearIncomingCtxMetadataAuthorizationToken clears the authorization token from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to clear the token from
//
// Returns:
//
//   - context.Context: The context with the token cleared
//   - error: An error if the metadata is not found or any other error occurs
func ClearIncomingCtxMetadataAuthorizationToken(ctx context.Context) (
	context.Context,
	error,
) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = ClearMetadataAuthorizationToken(md)
	return metadata.NewIncomingContext(ctx, md), nil
}

// ClearIncomingCtxMetadataGCloudAuthorizationToken clears the GCloud authorization token from the incoming context
// metadata
//
// Parameters:
//
//   - ctx: The incoming context to clear the token from
//
// Returns:
//
//   - context.Context: The context with the token cleared
//   - error: An error if the metadata is not found or any other error occurs
func ClearIncomingCtxMetadataGCloudAuthorizationToken(ctx context.Context) (
	context.Context,
	error,
) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = ClearMetadataAuthorizationToken(md)
	return metadata.NewIncomingContext(ctx, md), nil
}

// ClearIncomingCtxMetadataRefreshToken clears the refresh token from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to clear the token from
//
// Returns:
//
//   - context.Context: The context with the token cleared
//   - error: An error if the metadata is not found or any other error occurs
func ClearIncomingCtxMetadataRefreshToken(ctx context.Context) (
	context.Context,
	error,
) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = ClearMetadataRefreshToken(md)
	return metadata.NewIncomingContext(ctx, md), nil
}

// ClearIncomingCtxMetadataAccessToken clears the access token from the incoming context metadata
//
// Parameters:
//
//   - ctx: The incoming context to clear the token from
//
// Returns:
//
//   - context.Context: The context with the token cleared
//   - error: An error if the metadata is not found or any other error occurs
func ClearIncomingCtxMetadataAccessToken(ctx context.Context) (context.Context, error) {
	// Get the metadata from the context
	md, err := GetIncomingCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = ClearMetadataAccessToken(md)
	return metadata.NewIncomingContext(ctx, md), nil
}
