package metadata

import (
	"strings"

	gogrpc "github.com/ralvarezdev/go-grpc"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	"google.golang.org/grpc/metadata"
)

// GetMetadataValue gets the value for a given key from the metadata
//
// Parameters:
//
//   - md: The metadata to get the value from
//   - key: The key to get the value for
//
// Returns:
//
//   - []string: The value for the given key
//   - error: An error if the key is not found or any other error occurs
func GetMetadataValue(md metadata.MD, key string) ([]string, error) {
	// Get the value from the metadata
	value := md.Get(key)
	if value == nil {
		return nil, ErrNilMetadataKeyValue
	}
	return value, nil
}

// DeleteMetadataValue deletes the value for a given key from the metadata
//
// Parameters:
//
//   - md: The metadata to delete the value from
//   - key: The key to delete the value for
//
// Returns:
//
//   - metadata.MD: The metadata with the value deleted
func DeleteMetadataValue(md metadata.MD, key string) metadata.MD {
	md.Delete(key)
	return md
}

// GetMetadataBearerToken gets the bearer token from the metadata
//
// Parameters:
//
//   - md: The metadata to get the token from
//   - key: The key to get the token for
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetMetadataBearerToken(md metadata.MD, key string) (string, error) {
	// Get the value from the metadata
	value, err := GetMetadataValue(md, key)
	if err != nil {
		return "", err
	}

	// Check if the authorization value is valid
	if len(value) <= gojwtgrpc.AuthorizationTokenIdx {
		return "", gojwtgrpc.ErrAuthorizationMetadataNotProvided
	}

	// Get the authorization value from the metadata
	authorizationValue := value[gojwtgrpc.AuthorizationTokenIdx]

	// Split the authorization value by space
	authorizationFields := strings.Split(authorizationValue, " ")

	// Check if the authorization value is valid
	if len(authorizationFields) != 2 || authorizationFields[0] != gojwt.BearerPrefix {
		return "", gojwtgrpc.ErrAuthorizationMetadataInvalid
	}

	return authorizationFields[1], nil
}

// GetMetadataAuthorizationToken gets the authorization token from the metadata
//
// Parameters:
//
//   - md: The metadata to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetMetadataAuthorizationToken(md metadata.MD) (string, error) {
	return GetMetadataBearerToken(md, gogrpc.AuthorizationMetadataKey)
}

// GetMetadataGCloudAuthorizationToken gets the GCloud authorization token from the metadata
//
// Parameters:
//
//   - md: The metadata to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetMetadataGCloudAuthorizationToken(md metadata.MD) (string, error) {
	return GetMetadataBearerToken(md, gogrpc.GCloudAuthorizationMetadataKey)
}

// GetMetadataRefreshToken gets the refresh token from the metadata
//
// Parameters:
//
//   - md: The metadata to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetMetadataRefreshToken(md metadata.MD) (string, error) {
	return GetMetadataBearerToken(md, gogrpc.RefreshTokenMetadataKey)
}

// SetMetadataBearerToken sets the authorization token to the metadata
//
// Parameters:
//
//   - md: The metadata to set the token to
//   - key: The metadata key where the token will be set
//   - token: The token to set
//
// Returns:
//
//   - metadata.MD: The metadata with the token set
func SetMetadataBearerToken(md metadata.MD, key, token string) metadata.MD {
	md.Set(key, gojwt.BearerPrefix+" "+token)
	return md
}

// SetMetadataAuthorizationToken sets the authorization token to the metadata
//
// Parameters:
//
//   - md: The metadata to set the token to
//   - token: The token to set
//
// Returns:
//
//   - metadata.MD: The metadata with the token set
func SetMetadataAuthorizationToken(md metadata.MD, token string) metadata.MD {
	return SetMetadataBearerToken(md, gogrpc.AuthorizationMetadataKey, token)
}

// SetMetadataGCloudAuthorizationToken sets the GCloud authorization token to the metadata
//
// Parameters:
//
//   - md: The metadata to set the token to
//   - token: The token to set
//
// Returns:
//
//   - metadata.MD: The metadata with the token set
func SetMetadataGCloudAuthorizationToken(
	md metadata.MD,
	token string,
) metadata.MD {
	return SetMetadataBearerToken(
		md,
		gogrpc.GCloudAuthorizationMetadataKey,
		token,
	)
}

// SetMetadataRefreshToken sets the refresh token to the metadata
//
// Parameters:
//
//   - md: The metadata to set the token to
//   - refreshToken: The token to set
//
// Returns:
//
//   - metadata.MD: The metadata with the token set
func SetMetadataRefreshToken(md metadata.MD, refreshToken string) metadata.MD {
	return SetMetadataBearerToken(
		md,
		gogrpc.RefreshTokenMetadataKey,
		refreshToken,
	)
}

// SetMetadataAccessToken sets the access token to the metadata
//
// Parameters:
//
//   - md: The metadata to set the token to
//   - accessToken: The token to set
//
// Returns:
//
//   - metadata.MD: The metadata with the token set
func SetMetadataAccessToken(md metadata.MD, accessToken string) metadata.MD {
	return SetMetadataBearerToken(
		md,
		gogrpc.AccessTokenMetadataKey,
		accessToken,
	)
}

// GetMetadataAccessToken gets the access token from the metadata
//
// Parameters:
//
//   - md: The metadata to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetMetadataAccessToken(md metadata.MD) (string, error) {
	return GetMetadataBearerToken(md, gogrpc.AccessTokenMetadataKey)
}

// ClearMetadataAuthorizationToken clears the authorization token from the metadata
//
// Parameters:
//
//   - md: The metadata to clear the token from
//
// Returns:
//
//   - metadata.MD: The metadata with the token cleared
func ClearMetadataAuthorizationToken(md metadata.MD) metadata.MD {
	return DeleteMetadataValue(md, gogrpc.AuthorizationMetadataKey)
}

// ClearMetadataGCloudAuthorizationToken clears the GCloud authorization token from the metadata
//
// Parameters:
//
//   - md: The metadata to clear the token from
//
// Returns:
//
//   - metadata.MD: The metadata with the token cleared
func ClearMetadataGCloudAuthorizationToken(md metadata.MD) metadata.MD {
	return DeleteMetadataValue(md, gogrpc.GCloudAuthorizationMetadataKey)
}

// ClearMetadataRefreshToken clears the refresh token from the metadata
//
// Parameters:
//
//   - md: The metadata to clear the token from
//
// Returns:
//
//   - metadata.MD: The metadata with the token cleared
func ClearMetadataRefreshToken(md metadata.MD) metadata.MD {
	return DeleteMetadataValue(md, gogrpc.RefreshTokenMetadataKey)
}

// ClearMetadataAccessToken clears the access token from the metadata
//
// Parameters:
//
//   - md: The metadata to clear the token from
//
// Returns:
//
//   - metadata.MD: The metadata with the token cleared
func ClearMetadataAccessToken(md metadata.MD) metadata.MD {
	return DeleteMetadataValue(md, gogrpc.AccessTokenMetadataKey)
}
