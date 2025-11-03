package gogrpc

const (
	// AuthorizationMetadataKey is the key used for authorization in metadata.
	AuthorizationMetadataKey = "authorization"

	// AuthorizationMetadataIndex is the index used for authorization in metadata.
	AuthorizationMetadataIndex = 0

	// RefreshTokenMetadataKey is the key used for refresh token in metadata.
	RefreshTokenMetadataKey = "x-refresh-token"

	// AccessTokenMetadataKey is the key used for access token in metadata.
	AccessTokenMetadataKey = "x-access-token"

	// GCloudAuthorizationMetadataKey is the key of the authorization metadata
	GCloudAuthorizationMetadataKey = "x-serverless-authorization"
)
