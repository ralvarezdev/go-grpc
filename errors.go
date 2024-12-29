package go_grpc

import (
	"net/http"
)

var (
	InternalServerError = http.StatusText(http.StatusInternalServerError)
)
