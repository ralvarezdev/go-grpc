package jwt

import (
	"context"
	"errors"

	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gogrpc "github.com/ralvarezdev/go-grpc"
	gogrpcmd "github.com/ralvarezdev/go-grpc/metadata"
	gogrpcserver "github.com/ralvarezdev/go-grpc/server"
)

type (
	// Interceptor is the interceptor for the authentication
	Interceptor struct {
		validator     gojwtvalidator.Validator
		interceptions map[string]*gojwttoken.Token
	}
)

// NewInterceptor creates a new authentication interceptor
//
// Parameters:
//
//   - validator: the JWT validator to validate the tokens
//   - interceptions: the gRPC interceptions to determine which methods require authentication
//
// Returns:
//
//   - *Interceptor: the interceptor
//   - error: if there was an error creating the interceptor
func NewInterceptor(
	validator gojwtvalidator.Validator,
	interceptions map[string]*gojwttoken.Token,
) (*Interceptor, error) {
	// Check if either the validator or the gRPC interceptions is nil
	if validator == nil {
		return nil, gojwtvalidator.ErrNilValidator
	}
	if interceptions == nil {
		return nil, gogrpc.ErrNilInterceptions
	}

	return &Interceptor{
		validator,
		interceptions,
	}, nil
}

// Authenticate returns the authentication interceptor
//
// Returns:
//
//   - grpc.UnaryServerInterceptor: the authentication interceptor
func (i Interceptor) Authenticate() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req any, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		// Check if the method should be intercepted
		interception, ok := i.interceptions[info.FullMethod]
		if !ok || interception == nil {
			return handler(ctx, req)
		}

		// Get the raw token from the metadata
		rawToken, err := gogrpcmd.GetIncomingCtxMetadataAuthorizationToken(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Validate the token and get the validated claims
		claims, err := i.validator.ValidateClaims(ctx, rawToken, *interception)
		if err != nil {
			if errors.Is(err, gojwtvalidator.ErrNilClaims) {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}

			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, status.Error(
					codes.Unauthenticated,
					gogrpcserver.ErrTokenHasExpired.Error(),
				)
			}

			return nil, status.Error(codes.Internal, gogrpc.InternalServerError)
		}

		// Set the raw token and token claims to the context
		ctx = gojwtgrpc.SetCtxToken(ctx, rawToken)
		ctx = gojwtgrpc.SetCtxTokenClaims(ctx, claims)

		return handler(ctx, req)
	}
}
