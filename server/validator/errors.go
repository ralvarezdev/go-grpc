package validator

import (
	"errors"
)

const (
	ErrFailedToCreateMapper   = "failed to create mapper on request: %v"
	ErrFailedToAssertValidationsToBadRequest = "failed to assert validations to bad request: %v"
	ErrFailedToValidateRequest = "failed to validate request: %v"
	ErrFailedToCreateValidateFunction = "failed to create validate function: %v"
)

var (
	ErrNilValidator = errors.New("validator is nil")
	ErrValidationsFailed = errors.New("validations failed")
)
