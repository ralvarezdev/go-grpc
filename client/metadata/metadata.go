package metadata

import (
	"context"
	"strings"

	gogrpcgcloud "github.com/ralvarezdev/go-grpc/cloud/gcloud"
	gojwt "github.com/ralvarezdev/go-jwt"
	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	"google.golang.org/grpc/metadata"
)

type (
	// Field is a field in the metadata
	Field struct {
		key   string
		value string
	}

	// CtxMetadata is the metadata for the context
	CtxMetadata []Field
)

// NewCtxMetadata creates a new CtxMetadata
//
// Parameters:
//
//   - metadataFields: The metadata fields to add to the context
//
// Returns:
//
//   - CtxMetadata: The context metadata
//   - error: An error if the metadata fields are nil or any other error occurs
func NewCtxMetadata(metadataFields map[string]string) (CtxMetadata, error) {
	// Check if the metadata fields are nil
	if metadataFields == nil {
		return nil, ErrNilMetadataFields
	}

	// Add the metadata fields
	var ctxMetadata CtxMetadata
	for key, value := range metadataFields {
		ctxMetadata = append(
			ctxMetadata,
			Field{key: strings.ToLower(key), value: value},
		)
	}

	return ctxMetadata, nil
}

// NewAuthenticatedCtxMetadata creates a new unauthenticated CtxMetadata
//
// Parameters:
//
//   - jwtToken: The JWT token to add to the context
//
// Returns:
//
//   - CtxMetadata: The context metadata
//   - error: An error if the JWT token is empty or any other error occurs
func NewAuthenticatedCtxMetadata(jwtToken string) (CtxMetadata, error) {
	return NewCtxMetadata(
		map[string]string{
			gojwtgrpc.AuthorizationMetadataKey: gojwt.BearerPrefix + " " + jwtToken,
		},
	)
}

// NewGCloudUnauthenticatedCtxMetadata creates a new unauthenticated CtxMetadata for GCloud
//
// Parameters:
//
//   - gcloudToken: The GCloud token to add to the context
//
// Returns:
//
//   - CtxMetadata: The context metadata
//   - error: An error if the GCloud token is empty or any other error occurs
func NewGCloudUnauthenticatedCtxMetadata(gcloudToken string) (
	CtxMetadata,
	error,
) {
	return NewCtxMetadata(
		map[string]string{
			gogrpcgcloud.AuthorizationMetadataKey: gojwt.BearerPrefix + " " + gcloudToken,
		},
	)
}

// NewGCloudAuthenticatedCtxMetadata creates a new authenticated CtxMetadata for GCloud
//
// Parameters:
//
//   - gcloudToken: The GCloud token to add to the context
//   - jwtToken: The JWT token to add to the context
//
// Returns:
//
//   - CtxMetadata: The context metadata
func NewGCloudAuthenticatedCtxMetadata(
	gcloudToken string, jwtToken string,
) (CtxMetadata, error) {
	return NewCtxMetadata(
		map[string]string{
			gogrpcgcloud.AuthorizationMetadataKey: gojwt.BearerPrefix + " " + gcloudToken,
			gojwtgrpc.AuthorizationMetadataKey:    gojwt.BearerPrefix + " " + jwtToken,
		},
	)
}

// GetCtxWithMetadata gets the context with the metadata
//
// Parameters:
//
//   - ctxMetadata: The context metadata to add to the context
//   - ctx: The context to add the metadata to
//
// Returns:
//
//   - context.Context: The context with the metadata
func GetCtxWithMetadata(
	ctxMetadata CtxMetadata, ctx context.Context,
) context.Context {
	// Check if the context metadata is nil
	if ctxMetadata == nil {
		return ctx
	}

	// Create metadata
	md := metadata.Pairs()

	// Add the metadata to the context
	for _, field := range ctxMetadata {
		md.Append(field.key, field.value)
	}
	return metadata.NewOutgoingContext(ctx, md)
}

// AppendGCloudTokenToOutgoingContext appends the GCloud token to the outgoing context
//
// Parameters:
//
//   - ctx: The context to append the GCloud token to
//   - gcloudToken: The GCloud token to append to the context
//
// Returns:
//
//   - context.Context: The context with the GCloud token appended
func AppendGCloudTokenToOutgoingContext(
	ctx context.Context, gcloudToken string,
) context.Context {
	return metadata.AppendToOutgoingContext(
		ctx,
		gogrpcgcloud.AuthorizationMetadataKey,
		gojwt.BearerPrefix+" "+gcloudToken,
	)
}
