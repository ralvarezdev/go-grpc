package metadata

import (
	"errors"
)

var (
	ErrNilMetadata         = errors.New("missing metadata")
	ErrNilMetadataKeyValue = errors.New("metadata key value is nil")
)
