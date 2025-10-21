package outgoing

import (
	"errors"
)

var (
	ErrFailedToGetOutgoingContext = errors.New(
		"failed to get outgoing context",
	)
)
