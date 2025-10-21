package errorhandler

import (
	"google.golang.org/grpc"
)

type (
	// ErrorHandler interface
	ErrorHandler interface {
		HandleError() grpc.UnaryServerInterceptor
	}
)
