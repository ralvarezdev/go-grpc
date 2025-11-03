package gogrpc

import (
	"errors"
	"net/http"
)

var (
	InternalServerError = http.StatusText(http.StatusInternalServerError)
)

var (
	ErrNilInterceptions = errors.New("grpc interceptions map cannot be nil")
)
