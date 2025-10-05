package auth

import (
	"google.golang.org/grpc"
)

type (
	// Authentication interface
	Authentication interface {
		Authenticate() grpc.UnaryClientInterceptor
	}
)
