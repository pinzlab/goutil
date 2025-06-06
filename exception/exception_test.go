package exception

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewExceptionWithAllFields(t *testing.T) {
	cause := errors.New("original cause")
	err := New("E123", "DatabaseError", cause)

	require.Error(t, err)

	ex, ok := err.(*Exception)
	require.True(t, ok, "error should be of type *Exception")

	assert.Equal(t, "E123", ex.Name)
	assert.Equal(t, "DatabaseError", ex.Description)
	assert.Equal(t, cause, ex.Cause)

	expectedMsg := "DatabaseError (E123): original cause"
	assert.Equal(t, expectedMsg, ex.Error())
}

func TestNewExceptionWithNoFields(t *testing.T) {
	err := New("", "", nil)

	require.Error(t, err)

	ex, ok := err.(*Exception)
	require.True(t, ok)

	assert.Empty(t, ex.Name)
	assert.Empty(t, ex.Description)
	assert.Nil(t, ex.Cause)

	assert.Equal(t, "exception", ex.Error())
}

func TestUnwrap(t *testing.T) {
	cause := errors.New("root error")
	err := New("CODE", "WrappedError", cause)

	unwrapped := errors.Unwrap(err)
	assert.Equal(t, cause, unwrapped)
}

func TestWithCauseAttachesError(t *testing.T) {
	baseErr := errors.New("failed to connect")
	ErrDatabase := Exception{
		Name:        "DB001",
		Description: "DatabaseError",
	}

	// Create a new error instance based on the template
	err := ErrDatabase.WithCause(baseErr)

	require.Error(t, err)

	ex, ok := err.(*Exception)
	require.True(t, ok, "should return *Exception")

	assert.Equal(t, "DB001", ex.Name)
	assert.Equal(t, "DatabaseError", ex.Description)
	assert.Equal(t, baseErr, ex.Cause)

	expectedMsg := "DatabaseError (DB001): failed to connect"
	assert.Equal(t, expectedMsg, err.Error())
}
func TestErrorsAs(t *testing.T) {
	cause := errors.New("something went wrong")
	err := New("ERR42", "BadInput", cause)

	var ex *Exception
	ok := errors.As(err, &ex)

	require.True(t, ok)
	assert.Equal(t, "ERR42", ex.Name)
	assert.Equal(t, "BadInput", ex.Description)
	assert.Equal(t, cause, ex.Cause)
}
