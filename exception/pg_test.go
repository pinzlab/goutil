package exception

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestException_ErrorMethod(t *testing.T) {
	ex := &Exception{
		Name:        "DATABASE_ERROR",
		Description: "Could not connect",
		Cause:       errors.New("timeout"),
	}

	expected := "Could not connect (DATABASE_ERROR): timeout"
	assert.Equal(t, expected, ex.Error())
}

func TestException_Error_NoDescription(t *testing.T) {
	ex := &Exception{
		Name:  "SOME_CODE",
		Cause: errors.New("internal"),
	}

	expected := "exception (SOME_CODE): internal"
	assert.Equal(t, expected, ex.Error())
}

func TestException_Error_NoCause(t *testing.T) {
	ex := &Exception{
		Name:        "NO_CAUSE",
		Description: "No underlying cause",
	}

	expected := "No underlying cause (NO_CAUSE)"
	assert.Equal(t, expected, ex.Error())
}

func TestException_Unwrap(t *testing.T) {
	orig := errors.New("original")
	ex := &Exception{
		Name:        "ERR",
		Description: "wrapped error",
		Cause:       orig,
	}

	assert.Equal(t, orig, errors.Unwrap(ex))
}

func TestException_WithCause(t *testing.T) {
	ex := Exception{
		Name:        "SIMPLE",
		Description: "Initial error",
	}

	cause := errors.New("root cause")
	newEx := ex.WithCause(cause).(*Exception)

	assert.Equal(t, "SIMPLE", newEx.Name)
	assert.Equal(t, "Initial error", newEx.Description)
	assert.Equal(t, cause, newEx.Cause)
}

func TestNew(t *testing.T) {
	cause := errors.New("db down")
	err := New("DB_ERR", "Database failure", cause)
	ex := err.(*Exception)

	assert.Equal(t, "DB_ERR", ex.Name)
	assert.Equal(t, "Database failure", ex.Description)
	assert.Equal(t, cause, ex.Cause)
}

func TestNewSimple(t *testing.T) {
	err := NewSimple("BAD_REQUEST", "Invalid input")
	ex := err.(*Exception)

	assert.Equal(t, "BAD_REQUEST", ex.Name)
	assert.Equal(t, "Invalid input", ex.Description)
	assert.Nil(t, ex.Cause)
}
