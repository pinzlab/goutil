package migrator

import (
	"fmt"
	"strings"
	"time"
)

// Entity represents a table and a set of rows to insert,
// skipping rows that already exist based on specified columns.
type Entity struct {
	Table   string   // Table name
	Check   []string // Columns to check for existing rows
	Columns []string // Columns to insert
	Values  [][]any  // Rows of values to insert
}

// GetScript returns a SQL script that inserts each row only if
// no existing row matches the Check columns.
func (e *Entity) GetScript() string {
	var result []string

	cols := map[string]int{}

	for index, col := range e.Columns {
		cols[col] = index
	}

	for _, row := range e.Values {
		var where []string
		var values []string

		for _, value := range row {
			switch any(value).(type) {
			case string:
				values = append(values, fmt.Sprintf("'%s'", value))
			case time.Time:
				formattedTime := (value).(time.Time).Format(time.RFC3339Nano)
				values = append(values, fmt.Sprintf("'%s'", formattedTime))
			case nil:
				values = append(values, "null")
			default:
				values = append(values, fmt.Sprintf("%+v", value))
			}
		}

		for _, check := range e.Check {
			where = append(where, fmt.Sprintf("%s = %s", check, values[cols[check]]))
		}

		result = append(result, fmt.Sprintf("INSERT INTO %s(%s) SELECT %s WHERE NOT EXISTS (SELECT 1 FROM %s WHERE %s);",
			e.Table,
			strings.Join(e.Columns, ", "),
			strings.Join(values, ", "),
			e.Table,
			strings.Join(where, " AND "),
		))

	}

	return strings.Join(result, "\n")

}
