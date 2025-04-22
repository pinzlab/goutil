package pg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetEnumScript tests the GetScript method of the Enum type.
func TestGetEnumScript(t *testing.T) {
	tests := []struct {
		name     string // name of the test case
		enum     Enum   // the Enum object to test
		expected string // the expected SQL script
	}{
		{
			name: "Simple enum",
			enum: Enum{
				Name:   "status",
				Values: []string{"active", "inactive", "pending"},
			},
			expected: `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
			CREATE TYPE status AS ENUM ('active', 'inactive', 'pending');
		END IF;
	END $$;
	`,
		},
		{
			name: "Enum with single value",
			enum: Enum{
				Name:   "level",
				Values: []string{"low"},
			},
			expected: `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'level') THEN
			CREATE TYPE level AS ENUM ('low');
		END IF;
	END $$;
	`,
		},
		{
			name: "Enum with special characters",
			enum: Enum{
				Name:   "priority",
				Values: []string{"high-priority", "medium", "low-priority"},
			},
			expected: `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'priority') THEN
			CREATE TYPE priority AS ENUM ('high-priority', 'medium', 'low-priority');
		END IF;
	END $$;
	`,
		},
	}

	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			script := item.enum.GetScript()
			assert.Equal(t, item.expected, script)
		})
	}
}
