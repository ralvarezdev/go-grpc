package validator

import (
	"time"

	"github.com/ralvarezdev/go-validator/struct/mapper/validation"
	govalidatorstructmappervalidator "github.com/ralvarezdev/go-validator/struct/mapper/validator"
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
			options *govalidatorstructmappervalidator.BirthdateOptions,
			validations *validation.StructValidations,
		)
		Password(
			passwordField string,
			password string,
			options *govalidatorstructmappervalidator.PasswordOptions,
			validations *validation.StructValidations,
		)
		CreateValidateFn(
			requestExample interface{},
			cache bool,
			auxiliaryValidatorFns ...govalidatorstructmappervalidator.AuxiliaryValidatorFn,
		) (ValidateFn, error)
		Validate(
			request interface{},
			auxiliaryValidatorFns ...govalidatorstructmappervalidator.AuxiliaryValidatorFn,
		) error
	}
)
