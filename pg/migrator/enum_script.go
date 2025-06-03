package migrator

import "strings"

// Enum represents a PostgreSQL ENUM type, with its name and possible values.
type Enum struct {
	Name   string   // The name of the ENUM type
	Values []string // The possible values of the ENUM type
}

// Generates a SQL script that creates the ENUM type in PostgreSQL if it does not already exist.
// The script uses a DO block to execute the SQL command conditionally.
func (e *Enum) GetScript() string {
	return `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = '` + e.Name + `') THEN
			CREATE TYPE ` + e.Name + ` AS ENUM ('` + strings.Join(e.Values, "', '") + `');
		END IF;
	END $$;
	`
}
