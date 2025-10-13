package auth_verifier

import (
	"google.golang.org/grpc"
)

type (
	// AuthenticationVerifier interface
	AuthenticationVerifier interface {
		VerifyAuthentication() grpc.UnaryClientInterceptor
	}
)
