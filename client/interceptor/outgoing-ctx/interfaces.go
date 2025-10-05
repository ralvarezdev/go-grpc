package outgoing_ctx

import (
	"google.golang.org/grpc"
)

type (
	// OutgoingCtx interface
	OutgoingCtx interface {
		PrintOutgoingCtx() grpc.UnaryClientInterceptor
	}
)
