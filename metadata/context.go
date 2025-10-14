package metadata

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// GetCtxMetadata gets the metadata from the context
//
// Parameters:
//
//   - ctx: The context to get the metadata from
//
// Returns:
//
//   - metadata.MD: The metadata from the context
//   - error: An error if the metadata is not found or any other error occurs
func GetCtxMetadata(ctx context.Context) (metadata.MD, error) {
	// Get the metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrNilMetadata
	}
	return md, nil
}

// GetCtxMetadataValue gets the value for a given key from the context metadata
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
func GetCtxMetadataValue(ctx context.Context, key string) ([]string, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	return GetMetadataValue(md, key)
}

// DeleteCtxMetadataValue deletes the value for a given key from the context metadata
//
// Parameters:
//
//   - ctx: The context to get the metadata from
//   - key: The key to delete the value for
//
// Returns:
//
//   - context.Context: The context with the value deleted
//   - error: An error if the key is not found or any other error occurs
func DeleteCtxMetadataValue(ctx context.Context, key string) (
	context.Context,
	error,
) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = DeleteMetadataValue(md, key)
	return metadata.NewIncomingContext(ctx, md), nil
}

// GetCtxMetadataBearerToken gets the bearer token from the context metadata
//
// Parameters:
//
//   - ctx: The context to get the metadata from
//   - key: The key to get the token for
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetCtxMetadataBearerToken(ctx context.Context, key string) (
	string,
	error,
) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataBearerToken(md, key)
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
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataAuthorizationToken(md)
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
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataGCloudAuthorizationToken(md)
}

// GetCtxMetadataRefreshToken gets the refresh token from the context metadata
//
// Parameters:
//
//   - ctx: The context to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetCtxMetadataRefreshToken(ctx context.Context) (string, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataRefreshToken(md)
}

// SetCtxMetadataBearerToken sets the authorization token to the metadata
//
// Parameters:
//
//   - ctx: The context to set the token to
//   - key: The metadata key where the token will be set
//   - token: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
//   - error: An error if the metadata is not found or any other error occurs
func SetCtxMetadataBearerToken(
	ctx context.Context,
	key, token string,
) (context.Context, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = SetMetadataBearerToken(md, key, token)
	return metadata.NewIncomingContext(ctx, md), nil
}

// SetCtxMetadataAuthorizationToken sets the authorization token to the metadata
//
// Parameters:
//
//   - ctx: The context to set the token to
//   - token: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
func SetCtxMetadataAuthorizationToken(
	ctx context.Context,
	token string,
) (context.Context, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = SetMetadataAuthorizationToken(md, token)
	return metadata.NewIncomingContext(ctx, md), nil
}

// SetCtxMetadataGCloudAuthorizationToken sets the GCloud authorization token to the metadata
//
// Parameters:
//
//   - ctx: The context to set the token to
//   - token: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
//   - error: An error if the metadata is not found or any other error occurs
func SetCtxMetadataGCloudAuthorizationToken(
	ctx context.Context,
	token string,
) (context.Context, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = SetMetadataGCloudAuthorizationToken(
		md,
		token,
	)
	return metadata.NewIncomingContext(ctx, md), nil
}

// SetCtxMetadataRefreshToken sets the refresh token to the metadata
//
// Parameters:
//
//   - ctx: The context to set the token to
//   - refreshToken: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
//   - error: An error if the metadata is not found or any other error occurs
func SetCtxMetadataRefreshToken(
	ctx context.Context,
	refreshToken string,
) (context.Context, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = SetMetadataRefreshToken(
		md,
		refreshToken,
	)
	return metadata.NewIncomingContext(ctx, md), nil
}

// SetCtxMetadataAccessToken sets the access token to the metadata
//
// Parameters:
//
//   - ctx: The context to set the token to
//   - accessToken: The token to set
//
// Returns:
//
//   - context.Context: The context with the token set
//   - error: An error if the metadata is not found or any other error occurs
func SetCtxMetadataAccessToken(
	ctx context.Context,
	accessToken string,
) (context.Context, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = SetMetadataAccessToken(
		md,
		accessToken,
	)
	return metadata.NewIncomingContext(ctx, md), nil
}

// GetCtxMetadataAccessToken gets the access token from the context metadata
//
// Parameters:
//
//   - ctx: The context to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetCtxMetadataAccessToken(ctx context.Context) (string, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return "", err
	}
	return GetMetadataAccessToken(md)
}

// ClearCtxMetadataAuthorizationToken clears the authorization token from the context metadata
//
// Parameters:
//
//   - ctx: The context to clear the token from
//
// Returns:
//
//   - context.Context: The context with the token cleared
//   - error: An error if the metadata is not found or any other error occurs
func ClearCtxMetadataAuthorizationToken(ctx context.Context) (
	context.Context,
	error,
) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = ClearMetadataAuthorizationToken(md)
	return metadata.NewIncomingContext(ctx, md), nil
}

// ClearCtxMetadataGCloudAuthorizationToken clears the GCloud authorization token from the context metadata
//
// Parameters:
//
//   - ctx: The context to clear the token from
//
// Returns:
//
//   - context.Context: The context with the token cleared
//   - error: An error if the metadata is not found or any other error occurs
func ClearCtxMetadataGCloudAuthorizationToken(ctx context.Context) (
	context.Context,
	error,
) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = ClearMetadataAuthorizationToken(md)
	return metadata.NewIncomingContext(ctx, md), nil
}

// ClearCtxMetadataRefreshToken clears the refresh token from the context metadata
//
// Parameters:
//
//   - ctx: The context to clear the token from
//
// Returns:
//
//   - context.Context: The context with the token cleared
//   - error: An error if the metadata is not found or any other error occurs
func ClearCtxMetadataRefreshToken(ctx context.Context) (
	context.Context,
	error,
) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = ClearMetadataRefreshToken(md)
	return metadata.NewIncomingContext(ctx, md), nil
}

// ClearCtxMetadataAccessToken clears the access token from the context metadata
//
// Parameters:
//
//   - ctx: The context to clear the token from
//
// Returns:
//
//   - context.Context: The context with the token cleared
//   - error: An error if the metadata is not found or any other error occurs
func ClearCtxMetadataAccessToken(ctx context.Context) (context.Context, error) {
	// Get the metadata from the context
	md, err := GetCtxMetadata(ctx)
	if err != nil {
		return nil, err
	}
	md = ClearMetadataAccessToken(md)
	return metadata.NewIncomingContext(ctx, md), nil
}
