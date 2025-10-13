package go_grpc

const (
	// AuthorizationMetadataKey is the key used for authorization in metadata.
	AuthorizationMetadataKey = "authorization"

	// RefreshTokenMetadataKey is the key used for refresh token in metadata.
	RefreshTokenMetadataKey = "x-refresh-token"

	// AccessTokenMetadataKey is the key used for access token in metadata.
	AccessTokenMetadataKey = "x-access-token"

	// GCloudAuthorizationMetadataKey is the key of the authorization metadata
	GCloudAuthorizationMetadataKey = "x-serverless-authorization"
)
