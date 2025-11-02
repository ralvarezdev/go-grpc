package http

import (
	"context"
	"net/http"

	gojwtnethttpctx "github.com/ralvarezdev/go-jwt/net/http/context"

	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
)

// SetOutgoingCtxMetadataAuthorizationToken is a helper function to set the authorization metadata to the outgoing
// context
//
// # Retrieves the authorization from the request context and sets it to the outgoing context
//
// Parameters:
//
//   - r: the request to get the authorization from and from which to get the context
//
// Returns:
//
//   - context.Context: the context with the authorization metadata
//   - error: an error if the request is nil
func SetOutgoingCtxMetadataAuthorizationToken(
	r *http.Request,
) (context.Context, error) {
	if r == nil {
		return nil, ErrNilRequest
	}

	// Get the authorization from the request context
	token, err := gojwtnethttpctx.GetCtxToken(r)
	if err != nil {
		return nil, err
	}

	// Set the authorization metadata to the outgoing context
	return gogrpcmd.SetOutgoingCtxMetadataAuthorizationToken(r.Context(), token)
}
