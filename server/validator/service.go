package validator

import (
	"log/slog"
	"reflect"
	"time"

	goreflect "github.com/ralvarezdev/go-reflect"
	govalidatormapper "github.com/ralvarezdev/go-validator/mapper"
	govalidatormapperparser "github.com/ralvarezdev/go-validator/mapper/parser"
	govalidatormapperparsergrpc "github.com/ralvarezdev/go-validator/mapper/parser/grpc"
	govalidatormappervalidation "github.com/ralvarezdev/go-validator/mapper/validation"
	govalidatormappervalidator "github.com/ralvarezdev/go-validator/mapper/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (

	// DefaultService is the default struct validator service
	DefaultService struct {
		generator   govalidatormapper.Generator
		service     govalidatormappervalidator.Service
		validateFns map[string]ValidateFn
		logger      *slog.Logger
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
	// Initialize the raw parser
	rawParser := govalidatormapperparser.NewDefaultRawParser(logger)

	// Initialize the end parser
	endParser := govalidatormapperparsergrpc.NewDefaultEndParser()

	// Initialize the validator
	validator := govalidatormappervalidator.NewDefaultValidator(logger)

	// Initialize the service
	service, err := govalidatormappervalidator.NewDefaultService(
		rawParser,
		endParser,
		validator,
		logger,
	)
	if err != nil {
		return nil, err
	}

	// Initialize the generator
	generator := govalidatormapper.NewProtobufGenerator(logger)

	// Create a logger for the service
	if logger != nil {
		logger = logger.With(
			slog.String("component", "grpc_validator_service"),
		)
	}

	return &DefaultService{
		service:     service,
		generator:   generator,
		logger:      logger,
		validateFns: make(map[string]ValidateFn),
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
	validations *govalidatormappervalidation.StructValidations,
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
	validations *govalidatormappervalidation.StructValidations,
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
	options *govalidatormappervalidator.BirthdateOptions,
	validations *govalidatormappervalidation.StructValidations,
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
	options *govalidatormappervalidator.PasswordOptions,
	validations *govalidatormappervalidation.StructValidations,
) {
	d.service.Password(
		passwordField,
		password,
		options,
		validations,
	)
}

// createMapper creates a mapper for a given struct
//
// Parameters:
//
//   - structInstance: the struct instance to create the mapper for
//
// Returns:
//
//   - *govalidatormapper.Mapper: the mapper
//   - error: if there was an error creating the mapper
func (d DefaultService) createMapper(
	structInstance interface{},
) (*govalidatormapper.Mapper, reflect.Type, error) {
	// Get the type of the request
	structInstanceType := goreflect.GetTypeOf(structInstance)

	// Dereference the request type if it is a pointer
	if structInstanceType.Kind() == reflect.Pointer {
		structInstanceType = structInstanceType.Elem()
	} else {
		structInstance = &structInstance
	}

	// Create the mapper
	mapper, err := d.generator.NewMapper(structInstance)
	if err != nil {
		if d.logger != nil {
			d.logger.Error(
				"Failed to create mapper",
				slog.String("type", structInstanceType.String()),
				slog.Any("error", err),
			)
		}
		return nil, structInstanceType, err
	}
	return mapper, structInstanceType, nil
}

// CreateValidateFn creates a validate function for a given request example
//
// Parameters:
//
//   - requestExample: an example of the request to validate
//   - cache: whether to cache the validate function or not
//   - auxiliaryValidatorFns: auxiliary validator functions to use in the validation
//
// Returns:
//
//   - ValidateFn: the validate function
//   - error: if there was an error creating the validate function
func (d DefaultService) CreateValidateFn(
	requestExample interface{},
	cache bool,
	auxiliaryValidatorFns ...interface{},
) (ValidateFn, error) {
	// Create the mapper
	mapper, requestType, err := d.createMapper(requestExample)
	if err != nil {
		return nil, err
	}

	// Check if the validate function is already cached
	if cache && d.validateFns != nil {
		if validateFn, ok := d.validateFns[goreflect.UniqueTypeReference(mapper.GetStructInstance())]; ok {
			return validateFn, nil
		}
	}

	// Create the inner validate function
	innerValidateFn, err := d.service.CreateValidateFn(
		mapper,
		cache,
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

	// Create the wrapped validate function
	validateFn := func(request interface{}) error {
		// Get a new instance of the body
		dest := goreflect.NewInstanceFromType(requestType)

		// Validate the request
		validations, err := innerValidateFn(dest)
		if err != nil {
			if d.logger != nil {
				d.logger.Error(
					"Failed to validate request",
					slog.String("type", requestType.String()),
					slog.Any("error", err),
				)
			}
			return status.Error(codes.Internal, "failed to validate request")
		}

		// Check if the error is nil and there are no validations
		if validations == nil {
			return nil
		}

		// Assert validations to BadRequest
		errorDetails, ok := validations.(*errdetails.BadRequest)
		if !ok {
			if d.logger != nil {
				d.logger.Error(
					"Failed to assert validations to BadRequest",
					slog.String("type", requestType.String()),
					slog.Any("validations", validations),
				)
			}
			return status.Error(codes.Internal, "failed to validate request")
		}

		// Create status with details
		st := status.New(codes.InvalidArgument, "validation failed")
		stWithDetails, _ := st.WithDetails(errorDetails)

		// Return error in your gRPC handler
		return stWithDetails.Err()
	}

	// Cache the validate function
	if cache {
		if d.validateFns == nil {
			d.validateFns = make(map[string]ValidateFn)
		}
		d.validateFns[goreflect.UniqueTypeReference(mapper.GetStructInstance())] = validateFn
	}
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
	auxiliaryValidatorFns ...interface{},
) error {
	// Create and cache the validate function
	validateFn, err := d.CreateValidateFn(
		request,
		true,
		auxiliaryValidatorFns...,
	)
	if err != nil {
		return status.Error(codes.Internal, "failed to validate request")
	}

	// Execute the validate function
	return validateFn(request)
}
