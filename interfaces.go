package go_grpc

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type (
	// ErrorDetailsGenerator interface for generating gRPC error details
	ErrorDetailsGenerator interface {
		NewFieldViolation(field, description string) *errdetails.BadRequest_FieldViolation
		NewSingleFieldViolation(field, description string) []*errdetails.BadRequest_FieldViolation
		NewBadRequest(violations []*errdetails.BadRequest_FieldViolation) *errdetails.BadRequest
		NewSingleBadRequest(field, description string) *errdetails.BadRequest
		NewStructSingleFieldBadRequest(structExample interface{}, field, description string) *errdetails.BadRequest
	}
)
