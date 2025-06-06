package exception

import (
	"fmt"
)

// Exception represents a structured error with a name, a description,
// and an optional underlying cause. It implements the error interface.
type Exception struct {
	// Name is a short identifier for the error, such as an error code.
	Name string

	// Description provides a human-readable explanation of the error.
	Description string

	// Cause holds the underlying error that caused this exception, if any.
	Cause error
}

// Error implements the error interface.
// It returns a formatted string representation of the exception,
// including the description, name, and cause if they are present.
func (e *Exception) Error() string {
	base := "exception"
	if e.Description != "" {
		base = e.Description
	}
	if e.Name != "" {
		base = fmt.Sprintf("%s (%s)", base, e.Name)
	}
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", base, e.Cause)
	}
	return base
}

// Unwrap returns the underlying cause of the Exception,
// enabling support for errors.Unwrap, errors.Is, and errors.As.
func (e *Exception) Unwrap() error {
	return e.Cause
}

// WithCause returns a new Exception based on the current one but
// with a specified cause. This allows chaining errors while preserving
// the original error context.
func (e Exception) WithCause(cause error) error {
	return &Exception{
		Name:        e.Name,
		Description: e.Description,
		Cause:       cause,
	}
}

// New creates a new Exception with the given name, description, and cause.
// It returns the Exception as an error.
func New(name, description string, cause error) error {
	return &Exception{
		Name:        name,
		Description: description,
		Cause:       cause,
	}
}

// NewSimple creates a new Exception with the given name and description,
// omitting the cause. It's a convenience function for situations where
// there is no underlying error to wrap.
func NewSimple(name, description string) error {
	return &Exception{
		Name:        name,
		Description: description,
	}
}
