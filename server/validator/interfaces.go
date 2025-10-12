package validator

import (
	"time"

	"github.com/ralvarezdev/go-validator/mapper/validation"
)

type (
	// Service interface
	Service interface {
		Email(
			emailField string,
			email string,
			validations *validation.StructValidations,
		)
		Username(
			usernameField string,
			username string,
			validations *validation.StructValidations,
		)
		Birthdate(
			birthdateField string,
			birthdate time.Time,
			validations *validation.StructValidations,
		)
		Password(
			passwordField string,
			password string,
			validations *validation.StructValidations,
		)
		CreateValidateFn(
			requestExample interface{},
			cache bool,
			auxiliaryValidatorFns ...interface{},
		) (ValidateFn, error)
		Validate(
			request interface{},
			auxiliaryValidatorFns ...interface{},
		) error
	}
)
