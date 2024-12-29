package context

import (
	"errors"
)

var (
	ErrFailedToGetPeerFromContext = errors.New("failed to get peer from context")
)
