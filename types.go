package go_grpc

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

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
func NewFieldViolation(field, description string) *errdetails.BadRequest_FieldViolation {
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
func NewSingleFieldViolation(field, description string) []*errdetails.BadRequest_FieldViolation {
	return []*errdetails.BadRequest_FieldViolation{
		NewFieldViolation(field, description),
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
func NewBadRequest(violations []*errdetails.BadRequest_FieldViolation) *errdetails.BadRequest {
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
func NewSingleBadRequest(field, description string) *errdetails.BadRequest {
	return NewBadRequest(NewSingleFieldViolation(field, description))
}
