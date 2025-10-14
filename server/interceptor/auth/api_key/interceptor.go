package api_keys

import (
	"context"

	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	// Interceptor is the interceptor for API key authentication
	Interceptor struct {
		allowedKeys map[string]struct{}
	}
)

// NewInterceptor creates a new API key authentication interceptor
//
// Parameters:
//
//   - keys: the allowed API keys
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: if no API keys are provided
func NewInterceptor(keys []string) (*Interceptor, error) {
	// Check if no API keys are provided
	if len(keys) == 0 {

		return nil, ErrNoAPIKeysProvided
	}

	// Create a map of allowed API keys for efficient lookup
	allowed := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		allowed[k] = struct{}{}
	}
	return &Interceptor{allowedKeys: allowed}, nil
}

// Authenticate returns the API key authentication interceptor
//
// Returns:
//
//   - grpc.UnaryServerInterceptor: the interceptor
func (i *Interceptor) Authenticate() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Get the raw token from the metadata
		rawToken, err := gogrpcmd.GetCtxMetadataAuthorizationToken(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Validate the API key
		if _, allowed := i.allowedKeys[rawToken]; !allowed {
			return nil, status.Error(codes.Unauthenticated, "invalid API key")
		}
		return handler(ctx, req)
	}
}
