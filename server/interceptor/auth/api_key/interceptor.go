package api_keys

import (
	"context"

	goapikey "github.com/ralvarezdev/go-api-key"
	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	// Interceptor is the interceptor for API key authentication
	Interceptor struct {
		apiKeyService goapikey.BasicService
		interceptions map[string]struct{}
	}
)

// NewInterceptor creates a new API key authentication interceptor
//
// Parameters:
//
//   - apiKeyService: the API key basic service to validate the API keys
//   - methodsToIntercept: a slice of method names to intercept
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: if no API keys are provided
func NewInterceptor(
	apiKeyService goapikey.BasicService,
	methodsToIntercept []string,
) (
	*Interceptor,
	error,
) {
	// Check if the API key service is nil
	if apiKeyService == nil {
		return nil, goapikey.ErrNilService
	}

	// Create a map of methods to intercept for efficient lookup
	interceptions := make(map[string]struct{})
	if methodsToIntercept != nil && len(methodsToIntercept) != 0 {
		for _, method := range methodsToIntercept {
			interceptions[method] = struct{}{}
		}
	}

	return &Interceptor{
		apiKeyService: apiKeyService,
		interceptions: interceptions,
	}, nil
}

// Authenticate returns the API key authentication interceptor
//
// Returns:
//
//   - grpc.UnaryServerInterceptor: the interceptor
func (i Interceptor) Authenticate() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Check if the method should be intercepted
		_, ok := i.interceptions[info.FullMethod]
		if !ok {
			return handler(ctx, req)
		}

		// Get the raw token from the metadata
		rawToken, err := gogrpcmd.GetCtxMetadataAuthorizationToken(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Validate the API key
		if valid := i.apiKeyService.IsAPIKeyValid(rawToken); !valid {
			return nil, status.Error(codes.Unauthenticated, "invalid API key")
		}
		return handler(ctx, req)
	}
}
