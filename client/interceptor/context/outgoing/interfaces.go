package outgoing

import (
	"google.golang.org/grpc"
)

type (
	// OutgoingCtx interface
	OutgoingCtx interface {
		PrintOutgoingCtx() grpc.UnaryClientInterceptor
	}
)
