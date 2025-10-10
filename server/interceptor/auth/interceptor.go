package auth

import (
	"context"
	"errors"

	gogrpc "github.com/ralvarezdev/go-grpc"
	gogrpcserver "github.com/ralvarezdev/go-grpc/server"
	gogrpcservermd "github.com/ralvarezdev/go-grpc/server/metadata"
	gojwtgrpc "github.com/ralvarezdev/go-jwt/grpc"
	gojwtgrpcmd "github.com/ralvarezdev/go-jwt/grpc/context"
	gojwttoken "github.com/ralvarezdev/go-jwt/token"
	gojwtvalidator "github.com/ralvarezdev/go-jwt/token/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
		return nil, gojwtgrpc.ErrNilGRPCInterceptions
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
		ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Check if the method should be intercepted
		interception, ok := i.interceptions[info.FullMethod]
		if !ok || interception == nil {
			return handler(ctx, req)
		}

		// Get metadata from the context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(
				codes.Unauthenticated,
				gojwtgrpc.ErrMissingMetadata.Error(),
			)
		}

		// Get the raw token from the metadata
		rawToken, err := gogrpcservermd.GetAuthorizationTokenFromMetadata(md)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		// Validate the token and get the validated claims
		claims, err := i.validator.ValidateClaims(rawToken, *interception)
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
		ctx = gojwtgrpcmd.SetCtxRawToken(ctx, rawToken)
		ctx = gojwtgrpcmd.SetCtxTokenClaims(ctx, claims)

		return handler(ctx, req)
	}
}
