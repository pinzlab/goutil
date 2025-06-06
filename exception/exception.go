package exception

import (
	"fmt"
)

// Exception represents a custom error that includes optional
// metadata: a code, a name, and an underlying cause.
//
// All fields are optional and can be left empty or nil.
// This struct implements the standard error interface.
type Exception struct {
	// Code is an optional string representing a machine-readable error code.
	Code string

	// Name is an optional human-readable identifier for the type of error.
	Name string

	// Cause is the underlying error that triggered this exception, if any.
	Cause error
}

// Error returns a string representation of the Exception.
// It includes the name, code, and the cause (if present).
//
// This method satisfies the error interface.
func (e *Exception) Error() string {
	base := "exception"
	if e.Name != "" {
		base = e.Name
	}
	if e.Code != "" {
		base = fmt.Sprintf("%s (%s)", base, e.Code)
	}
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", base, e.Cause)
	}
	return base
}

// Unwrap returns the underlying cause of the Exception,
// allowing compatibility with errors.Unwrap and errors.Is/errors.As.
func (e *Exception) Unwrap() error {
	return e.Cause
}

// WithCause returns a copy of the exception with a dynamic cause attached.
func (e Exception) WithCause(cause error) error {
	return &Exception{
		Code:  e.Code,
		Name:  e.Name,
		Cause: cause,
	}
}

// New creates a new Exception instance with the given code, name, and cause.
//
// Parameters:
//   - code: a string representing the error code (optional).
//   - name: a string representing the error name (optional).
//   - cause: the underlying error (can be nil).
//
// Returns:
//   - An error that wraps the provided details.
func New(code, name string, cause error) error {
	return &Exception{
		Code:  code,
		Name:  name,
		Cause: cause,
	}
}
