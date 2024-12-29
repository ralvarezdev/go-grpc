package outgoing_ctx

import (
	"errors"
)

var (
	ErrFailedToGetOutgoingContext = errors.New(
		"failed to get outgoing context",
	)
)
