package metadata

import (
	gogrpcgcloud "github.com/ralvarezdev/go-grpc/cloud/gcloud"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

// GetTokenFromMetadata gets the token from the metadata
func GetTokenFromMetadata(md metadata.MD, tokenKey string) (string, error) {
	// Get the authorization from the metadata
	authorization := md.Get(tokenKey)
	if len(authorization) <= gojwtgrpc.AuthorizationTokenIdx {
		return "", gojwtgrpc.ErrAuthorizationMetadataNotProvided
	}

	// Get the authorization value from the metadata
	authorizationValue := authorization[gojwtgrpc.AuthorizationTokenIdx]

	// Split the authorization value by space
	authorizationFields := strings.Split(authorizationValue, " ")

	// Check if the authorization value is valid
	if len(authorizationFields) != 2 || authorizationFields[0] != gojwt.BearerPrefix {
		return "", gojwtgrpc.ErrAuthorizationMetadataInvalid
	}

	return authorizationFields[1], nil
}

// GetAuthorizationTokenFromMetadata gets the authorization token from the metadata
func GetAuthorizationTokenFromMetadata(md metadata.MD) (string, error) {
	return GetTokenFromMetadata(md, gojwtgrpc.AuthorizationMetadataKey)
}

// GetGCloudAuthorizationTokenFromMetadata gets the GCloud authorization token from the metadata
func GetGCloudAuthorizationTokenFromMetadata(md metadata.MD) (string, error) {
	return GetTokenFromMetadata(md, gogrpcgcloud.AuthorizationMetadataKey)
}
