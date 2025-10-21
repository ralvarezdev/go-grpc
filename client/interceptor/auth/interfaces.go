package auth

import (
	"google.golang.org/grpc"
)

type (
	// Authenticator interface
	Authenticator interface {
		Authenticate() grpc.UnaryClientInterceptor
	}

	// Verifier interface
	Verifier interface {
		Verify() grpc.UnaryClientInterceptor
	}
)
