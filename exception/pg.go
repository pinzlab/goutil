package exception

import (
	"errors"
	"sync"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// PGError defines a mapping between PostgreSQL error codes or constraint names
// and user-friendly error messages.
//
// Example:
//
//	var customCodes = PGError{
//	    "23505": "Email already exists",
//	}
type PGError = map[string]string

// pgErrorCodes maps PostgreSQL error codes to default error messages.
var pgErrorCodes = PGError{
	"default":   "Query error",
	"not_found": "Record not found",
}

// pgConstraintError maps specific constraint names to friendly messages.
//
// Example:
//
//	"users_email_key" → "Email must be unique"
var pgConstraintError = PGError{}

// mu ensures thread-safe access to the error maps.
var mu sync.RWMutex

// InitPGError allows overriding or extending the default PostgreSQL error messages
// and constraint messages.
//
// You can call this once at startup to customize messages.
//
// Example:
//
//	func init() {
//	    customCodes := exception.PGError{
//	        "default":   "Query error",
//	        "not_found": "Record not found",
//	        "23505": "Custom duplicate error",
//	    }
//	    constraintMessages := exception.PGError{
//	        "users_email_key": "Email already exists",
//	    }
//	    exception.InitPGError(&customCodes, &constraintMessages)
//	}
func InitPGError(errorCodes, constraintCodes *PGError) {
	mu.Lock()
	defer mu.Unlock()

	if errorCodes != nil {
		for k, v := range *errorCodes {
			pgErrorCodes[k] = v
		}
	}

	if constraintCodes != nil {
		for k, v := range *constraintCodes {
			pgConstraintError[k] = v
		}
	}
}

// PG converts a low-level database or ORM error into a user-friendly Exception.
//
// It handles:
//   - PostgreSQL errors (*pgconn.PgError)
//   - GORM not found errors (gorm.ErrRecordNotFound)
//   - All other error types (as generic query error)
//
// Example:
//
//	err := &pgconn.PgError{
//	    Code:           "23505",
//	    ConstraintName: "users_email_key",
//	}
//
//	ex := exception.PG(err).(*exception.Exception)
//	fmt.Println(ex.Name)        // Output: "Duplicate record"
//	fmt.Println(ex.Description) // Output: "Email already exists"
func PG(err error) error {
	mu.RLock()
	defer mu.RUnlock()

	// PostgreSQL-specific error
	if pgErr, ok := err.(*pgconn.PgError); ok {
		return &Exception{
			Cause:       pgErr,
			Name:        getOrDefault(pgErrorCodes, pgErr.Code, pgErrorCodes["default"]),
			Description: getOrDefault(pgConstraintError, pgErr.ConstraintName, pgErr.Message),
		}
	}

	// Record not found (GORM)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Exception{
			Name:  pgErrorCodes["not_found"],
			Cause: err,
		}
	}

	// Generic or unknown error
	return &Exception{
		Name:  pgErrorCodes["default"],
		Cause: err,
	}
}

// getOrDefault safely retrieves a value from a map.
// If the key is not found, it returns the fallback value.
//
// Example:
//
//	getOrDefault(map[string]string{"foo": "bar"}, "foo", "default") → "bar"
//	getOrDefault(map[string]string{}, "baz", "default") → "default"
func getOrDefault(m map[string]string, key, fallback string) string {
	if val, ok := m[key]; ok {
		return val
	}
	return fallback
}
