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
			requestExample any,
			cache bool,
			auxiliaryValidatorFns ...any,
		) (ValidateFn, error)
		Validate(
			request any,
			auxiliaryValidatorFns ...any,
		) error
	}
)
