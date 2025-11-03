package metadata

import (
	"errors"
)

var (
	ErrNilMetadata                      = errors.New("missing metadata")
	ErrNilMetadataKeyValue              = errors.New("metadata key value is nil")
	ErrAuthorizationMetadataInvalid     = errors.New("authorization metadata invalid")
	ErrAuthorizationMetadataNotProvided = errors.New("authorization metadata is not provided")
)
