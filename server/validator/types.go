package validator

type (
	// ValidateFn func type for validating a value
	ValidateFn func(request any) error
)
