package http

import (
	"context"
	"net/http"

	gogrpcclientmd "github.com/ralvarezdev/go-grpc/client/metadata"
	gojwtnethttpctx "github.com/ralvarezdev/go-jwt/net/http/context"
)

// SetCtxAuthorizationMetadata is a helper function to set the authorization metadata in the context
//
// Retrieves the authorization from the request context and sets it in the provided context.
//
// Parameters:
//
//   - ctx: the context
//   - r: the request
//
// Returns:
//
//   - context.Context: the context with the authorization metadata
//   - error: an error if the request is nil
func SetCtxAuthorizationMetadata(
	ctx context.Context,
	r *http.Request,
) (context.Context, error) {
	if r == nil {
		return ctx, ErrNilRequest
	}

	// Get the authorization from the request context
	token, err := gojwtnethttpctx.GetCtxToken(r)
	if err != nil {
		return nil, err
	}

	// Set the authorization metadata in the context
	return gogrpcclientmd.SetCtxAuthorization(ctx, token), nil
}
