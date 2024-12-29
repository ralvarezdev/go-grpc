package server

import (
	"errors"
)

var (
	ErrTokenHasExpired = errors.New("token has expired")
)
