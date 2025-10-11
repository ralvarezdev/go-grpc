package metadata

import (
	"strings"

	gogrpcgcloud "github.com/ralvarezdev/go-grpc/cloud/gcloud"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	"google.golang.org/grpc/metadata"
)

// GetValueFromMetadata gets the value for a given key from the metadata
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
func GetValueFromMetadata(md metadata.MD, key string) ([]string, error) {
	// Get the value from the metadata
	value := md.Get(key)
	if value == nil {
		return nil, ErrNilMetadataKeyValue
	}
	return value, nil
}

// GetTokenFromMetadata gets the token from the metadata for a given key
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
func GetTokenFromMetadata(md metadata.MD, key string) (string, error) {
	// Get the value from the metadata
	value, err := GetValueFromMetadata(md, key)
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

// GetAuthorizationTokenFromMetadata gets the authorization token from the metadata
//
// Parameters:
//
//   - md: The metadata to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetAuthorizationTokenFromMetadata(md metadata.MD) (string, error) {
	return GetTokenFromMetadata(md, gojwtgrpc.AuthorizationMetadataKey)
}

// GetGCloudAuthorizationTokenFromMetadata gets the GCloud authorization token from the metadata
//
// Parameters:
//
//   - md: The metadata to get the token from
//
// Returns:
//
//   - string: The token
//   - error: An error if the token is not found or any other error occurs
func GetGCloudAuthorizationTokenFromMetadata(md metadata.MD) (string, error) {
	return GetTokenFromMetadata(md, gogrpcgcloud.AuthorizationMetadataKey)
}
