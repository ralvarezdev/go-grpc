package apikey

import (
	"errors"
)

var (
	ErrEmptyAPIKey                           = errors.New("empty API key")
	ErrFailedToSetMetadataAuthorizationToken = errors.New("failed to set metadata authorization token")
)
