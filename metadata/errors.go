package metadata

import (
	"errors"
)

var (
	ErrNilMetadata         = errors.New("metadata is nil")
	ErrNilMetadataKeyValue = errors.New("metadata key value is nil")
)
