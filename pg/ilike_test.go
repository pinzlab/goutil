package pg

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIlike(t *testing.T) {
	tests := []struct {
		value    string
		columns  []string
		expected Ilike
	}{
		{
			value:   "test",
			columns: []string{"column1", "column2"},
			expected: Ilike{
				Where: "UNACCENT(column1) ILIKE UNACCENT(@key) OR UNACCENT(column2) ILIKE UNACCENT(@key)",
				Args:  sql.Named("key", "%test%"),
			},
		},
		{
			value:   "search",
			columns: []string{"name", "description"},
			expected: Ilike{
				Where: "UNACCENT(name) ILIKE UNACCENT(@key) OR UNACCENT(description) ILIKE UNACCENT(@key)",
				Args:  sql.Named("key", "%search%"),
			},
		},
		{
			value:    "empty",
			columns:  []string{},
			expected: Ilike{},
		},
		{
			value:   "value",
			columns: []string{"singleColumn"},
			expected: Ilike{
				Where: "UNACCENT(singleColumn) ILIKE UNACCENT(@key)",
				Args:  sql.Named("key", "%value%"),
			},
		},
	}

	for _, item := range tests {

		t.Run(item.value, func(t *testing.T) {
			result := NewIlike(item.value, item.columns...)
			assert.Equal(t, item.expected.Where, result.Where, "Where clause mismatch")
			assert.Equal(t, item.expected.Args, result.Args, "Args mismatch")
		})
	}
}
