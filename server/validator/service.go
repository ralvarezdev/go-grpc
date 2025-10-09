package validator

import (
	"encoding/json"
	"log/slog"
	"reflect"
	"time"

	goreflect "github.com/ralvarezdev/go-reflect"
	govalidatorstructmapper "github.com/ralvarezdev/go-validator/struct/mapper"
	govalidatorstructmapperparser "github.com/ralvarezdev/go-validator/struct/mapper/parser"
	govalidatorstructmapperparserjson "github.com/ralvarezdev/go-validator/struct/mapper/parser/json"
	govalidatorstructmappervalidation "github.com/ralvarezdev/go-validator/struct/mapper/validation"
	govalidatorstructmappervalidator "github.com/ralvarezdev/go-validator/struct/mapper/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type (

	// DefaultService is the default struct validator service
	DefaultService struct {
		generator        govalidatorstructmapper.Generator
		parser           govalidatorstructmapperparser.Parser
		validator        govalidatorstructmappervalidator.Validator
		service          govalidatorstructmappervalidator.Service
		cacheValidateFns map[string]ValidateFn
		logger           *slog.Logger
	}
)

// NewService creates a new validator service
//
// Parameters:
//
//   - service: the validator service
//   - logger: the logger
//
// Returns:
//
//   - *Validator: the validator
//   - error: if there was an error creating the validator service
func NewService(
	logger *slog.Logger,
) (*DefaultService, error) {
	// Initialize the parser
	parser := govalidatorstructmapperparserjson.NewParser(logger)

	// Initialize the validator
	validator := govalidatorstructmappervalidator.NewDefaultValidator(logger)

	// Initialize the service
	service, err := govalidatorstructmappervalidator.NewDefaultService(
		parser,
		validator,
	)
	if err != nil {
		return nil, err
	}

	// Initialize the generator
	generator := govalidatorstructmapper.NewProtobufGenerator(logger)

	// Create a logger for the service
	if logger != nil {
		logger = logger.With(
			slog.String("component", "validator_service"),
		)
	}

	return &DefaultService{
		parser:           parser,
		validator:        validator,
		service:          service,
		generator:        generator,
		logger:           logger,
		cacheValidateFns: make(map[string]ValidateFn),
	}, nil
}

// Email validates an email field
//
// Parameters:
//
//   - emailField: the name of the email field
//   - email: the email to validate
//   - validations: the struct validations
func (d DefaultService) Email(
	emailField string,
	email string,
	validations *govalidatorstructmappervalidation.StructValidations,
) {
	d.service.Email(
		emailField,
		email,
		validations,
	)
}

// Username validates a username field
//
// Parameters:
//
//   - usernameField: the name of the username field
//   - username: the username to validate
//   - validations: the struct validations
func (d DefaultService) Username(
	usernameField string,
	username string,
	validations *govalidatorstructmappervalidation.StructValidations,
) {
	d.service.Username(
		usernameField,
		username,
		validations,
	)
}

// Birthdate validates a birthdate field
//
// Parameters:
//
//   - birthdateField: the name of the birthdate field
//   - birthdate: the birthdate to validate
//   - options: the birthdate options
//   - validations: the struct validations
func (d DefaultService) Birthdate(
	birthdateField string,
	birthdate time.Time,
	options *govalidatorstructmappervalidator.BirthdateOptions,
	validations *govalidatorstructmappervalidation.StructValidations,
) {
	d.service.Birthdate(
		birthdateField,
		birthdate,
		options,
		validations,
	)
}

// Password validates a password field
//
// Parameters:
//
//   - passwordField: the name of the password field
//   - password: the password to validate
//   - options: the password options
//   - validations: the struct validations
func (d DefaultService) Password(
	passwordField string,
	password string,
	options *govalidatorstructmappervalidator.PasswordOptions,
	validations *govalidatorstructmappervalidation.StructValidations,
) {
	d.service.Password(
		passwordField,
		password,
		options,
		validations,
	)
}

// CreateValidateFn creates a validate function for a given request example
//
// Parameters:
//
//   - requestExample: an example of the request to validate
//   - auxiliaryValidatorFns: auxiliary validator functions to use in the validation
//
// Returns:
//
//   - ValidateFn: the validate function
//   - error: if there was an error creating the validate function
func (d DefaultService) CreateValidateFn(
	requestExample interface{},
	auxiliaryValidatorFns ...govalidatorstructmappervalidator.AuxiliaryValidatorFn,
) (ValidateFn, error) {
	// Get the type of the request
	requestType := goreflect.GetTypeOf(requestExample)

	// Dereference the request type if it is a pointer
	if requestType.Kind() == reflect.Pointer {
		requestType = requestType.Elem()
	} else {
		requestExample = &requestExample
	}

	// Create the mapper
	mapper, err := d.generator.NewMapper(requestExample)
	if err != nil {
		if d.logger != nil {
			d.logger.Error(
				"Failed to create mapper",
				slog.String("type", requestType.String()),
				slog.Any("error", err),
			)
		}
		return nil, err
	}

	// Create the validate function
	validateFn, err := d.service.CreateValidateFn(
		mapper,
		auxiliaryValidatorFns...,
	)
	if err != nil {
		if d.logger != nil {
			d.logger.Error(
				"Failed to create validate function",
				slog.String("type", requestType.String()),
				slog.Any("error", err),
			)
		}
		return nil, err
	}

	return func(request interface{}) error {
		// Get a new instance of the body
		dest := goreflect.NewInstanceFromType(requestType)

		// Validate the request
		validations, err := validateFn(dest)
		if err != nil {
			if d.logger != nil {
				d.logger.Error(
					"Failed to validate request",
					slog.String("type", requestType.String()),
					slog.Any("error", err),
				)
			}
			return status.Error(codes.Internal, "Failed to validate request")
		}

		// Check if the error is nil and there are no validations
		if validations == nil {
			return nil
		}

		// Marshal to JSON
		jsonBytes, _ := json.Marshal(validations)
		jsonString := string(jsonBytes)

		// Wrap JSON string in StringValue
		stringValue := wrapperspb.StringValue{Value: jsonString}

		// Marshal to Any
		anyValue, _ := anypb.New(&stringValue)

		// Create status with details
		st := status.New(codes.InvalidArgument, "Validation failed")
		stWithDetails, _ := st.WithDetails(anyValue)

		// Return error in your gRPC handler
		return stWithDetails.Err()
	}, nil
}

// CreateAndCacheValidateFn creates and caches a validate function for a given request example
//
// Parameters:
//
//   - requestExample: an example of the request to validate
//   - auxiliaryValidatorFns: auxiliary validator functions to use in the validation
//
// Returns:
//
//   - ValidateFn: the validate function
//   - error: if there was an error creating the validate function
func (d DefaultService) CreateAndCacheValidateFn(
	requestExample interface{},
	auxiliaryValidatorFns ...govalidatorstructmappervalidator.AuxiliaryValidatorFn,
) (ValidateFn, error) {
	// Get the type of the request
	requestType := goreflect.GetTypeOf(requestExample)

	// Dereference the request type if it is a pointer
	if requestType.Kind() == reflect.Pointer {
		requestType = requestType.Elem()
	} else {
		requestExample = &requestExample
	}

	// Get the unique string representation of the request type
	uniqueReference := goreflect.UniqueTypeReference(requestType)

	// Check if the validate function is already cached
	if validateFn, ok := d.cacheValidateFns[uniqueReference]; ok {
		return validateFn, nil
	}

	// Create the validate function
	validateFn, err := d.CreateValidateFn(
		requestExample,
		auxiliaryValidatorFns...,
	)
	if err != nil {
		return nil, err
	}

	// Cache the validate function
	d.cacheValidateFns[uniqueReference] = validateFn

	return validateFn, nil
}

// Validate is the function that creates (if not cached), caches and executes the validation
//
// Parameters:
//
//   - request: the request to validate
//   - auxiliaryValidatorFns: auxiliary validator functions to use in the validation
//
// Returns:
//
//   - error: if there was an error validating the request
func (d DefaultService) Validate(
	request interface{},
	auxiliaryValidatorFns ...govalidatorstructmappervalidator.AuxiliaryValidatorFn,
) error {
	// Create and cache the validate function
	validateFn, err := d.CreateAndCacheValidateFn(
		request,
		auxiliaryValidatorFns...,
	)
	if err != nil {
		return status.Error(codes.Internal, "Failed to validate request")
	}

	// Execute the validate function
	return validateFn(request)
}
