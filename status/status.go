package status

import (
	"errors"
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gogrpc "github.com/ralvarezdev/go-grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ExtractErrorFromStatus extracts the error from the status
func ExtractErrorFromStatus(mode *goflagsmode.Flag, err error) (
	codes.Code,
	error,
) {
	st, ok := status.FromError(err)

	// Check if the error is a status error
	if !ok {
		// Check the flag mode
		if mode != nil && mode.IsProd() {
			return codes.Internal, errors.New(gogrpc.InternalServerError)
		}
		return codes.Internal, err
	}

	// Check the code
	code := st.Code()

	return code, errors.New(st.Message())
}
