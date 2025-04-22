package pg

import (
	"database/sql"
	"fmt"
	"strings"
)

// Ilike struct represents a query condition for case-insensitive matching
// using the ILIKE operator and a named argument for the value to be searched.
type Ilike struct {
	Where string       // SQL WHERE clause for the ILIKE condition
	Args  sql.NamedArg // Argument for the query with the search term
}

// NewIlike constructs an Ilike struct for a case-insensitive search.
// It generates a WHERE clause that matches the given value against
// the specified columns using the ILIKE operator with unaccented values.
//
// Parameters:
//   - value: The search term to match against the columns
//   - columns: The names of the columns to search within
//
// Returns:
//   - An Ilike struct containing the WHERE clause and argument for the search
func NewIlike(value string, columns ...string) Ilike {

	var unaccent Ilike

	if len(columns) == 0 {
		return unaccent
	}

	for index, col := range columns {
		columns[index] = fmt.Sprintf("UNACCENT(%s) ILIKE UNACCENT(@key)", col)
	}

	unaccent.Where = strings.Join(columns, " OR ")
	unaccent.Args = sql.Named("key", "%"+value+"%")

	return unaccent
}
