package go_grpc

import (
	"log/slog"

	goreflect "github.com/ralvarezdev/go-reflect"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type (
	// DefaultErrorDetailsGenerator is the default implementation of ErrorDetailsGenerator
	DefaultErrorDetailsGenerator struct {
		logger *slog.Logger
	}
)

// NewDefaultErrorDetailsGenerator creates a new DefaultErrorDetailsGenerator
//
// Parameters:
//
//   - logger: the logger (optional, can be nil)
//
// Returns:
//
//   - *DefaultErrorDetailsGenerator: the created DefaultErrorDetailsGenerator
func NewDefaultErrorDetailsGenerator(logger *slog.Logger) *DefaultErrorDetailsGenerator {
	if logger != nil {
		logger = logger.With(
			slog.String("component", "grpc_error_details_generator"),
		)
	}

	return &DefaultErrorDetailsGenerator{
		logger: logger,
	}
}

// NewFieldViolation creates a new field violation
//
// Parameters:
//
//   - field: the field that caused the violation
//   - description: a description of the violation
//
// Returns:
//
//   - *errdetails.BadRequest_FieldViolation: the created field violation
func (d DefaultErrorDetailsGenerator) NewFieldViolation(field, description string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: description,
	}
}

// NewSingleFieldViolation creates a new bad request with a single field violation
//
// Parameters:
//
//   - field: the field that caused the violation
//   - description: a description of the violation
//
// Returns:
//
//   - []*errdetails.BadRequest_FieldViolation: the created field violations
func (d DefaultErrorDetailsGenerator) NewSingleFieldViolation(field, description string) []*errdetails.BadRequest_FieldViolation {
	return []*errdetails.BadRequest_FieldViolation{
		d.NewFieldViolation(field, description),
	}
}

// NewBadRequest creates a new bad request with the given field violations
//
// Parameters:
//
//   - violations: the field violations
//
// Returns:
//
//   - *errdetails.BadRequest: the created bad request
func (d DefaultErrorDetailsGenerator) NewBadRequest(violations []*errdetails.BadRequest_FieldViolation) *errdetails.BadRequest {
	return &errdetails.BadRequest{
		FieldViolations: violations,
	}
}

// NewSingleBadRequest creates a new bad request with a single field violation
//
// Parameters:
//
//   - field: the field that caused the violation
//   - description: a description of the violation
//
// Returns:
//
//   - *errdetails.BadRequest: the created bad request
func (d DefaultErrorDetailsGenerator) NewSingleBadRequest(field, description string) *errdetails.BadRequest {
	return d.NewBadRequest(d.NewSingleFieldViolation(field, description))
}

// NewStructSingleFieldBadRequest creates a new bad request with a single field violation for a struct field
//
// Parameters:
//
//   - structExample: the struct example
//   - field: the field that caused the violation
//   - description: a description of the violation
//
// Returns:
//
//   - *errdetails.BadRequest: the created bad request
func (d DefaultErrorDetailsGenerator) NewStructSingleFieldBadRequest(
	structExample interface{},
	field, description string,
) *errdetails.BadRequest {
	// Warn if structExample is nil
	if structExample == nil && d.logger != nil {
		d.logger.Warn(
			"structExample is nil, cannot get struct type name for field violation",
			slog.String("field", field),
			slog.String("description", description),
		)
		return d.NewSingleBadRequest(field, description)
	}

	// Warn if the struct doesn't have the field
	reflection := goreflect.NewDereferencedReflection(structExample)
	if !reflection.HasField(field) {
		if d.logger != nil {
			d.logger.Warn(
				"structExample does not have field, cannot get struct type name for field violation",
				slog.String("field", field),
				slog.String("description", description),
			)
		}
		return d.NewSingleBadRequest(field, description)
	}

	return d.NewSingleBadRequest(field, description)
}
