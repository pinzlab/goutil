package migrator

import "strings"

// Unique represents a unique index on a table with specific columns.
// The index is applied only to records that are not logically deleted (WHERE dat IS NULL).
type Unique struct {
	Table   string   // Name of the table
	Columns []string // List of columns to apply the unique constraint
}

// GetScript generates the SQL script to create a unique index for the table and columns.
// The index is applied only to records where the 'dat' column is NULL, indicating the record is not logically deleted.
func (u *Unique) GetScript() string {
	return `
	CREATE UNIQUE INDEX IF NOT EXISTS uni_` + u.Table + `_` + strings.Join(u.Columns, "_") + `
		ON public.` + u.Table + `(` + strings.Join(u.Columns, ", ") + `)
	WHERE dat IS NULL;
	`
}
