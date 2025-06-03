package migrator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInsert tests the Insert method of the table record.
func TestInsert(t *testing.T) {

	tests := []struct {
		name     string // name of the test case
		entity   Entity // the Entity object to test
		expected string // the expected SQL script
	}{
		{
			name: "Simple insert",
			entity: Entity{
				Table:   "product",
				Check:   []string{"name"},
				Columns: []string{"name", "category", "description"},
				Values:  [][]any{{"mobile", nil, "Apple"}},
			},
			expected: "INSERT INTO product(name, category, description) SELECT 'mobile', null, 'Apple' " +
				"WHERE NOT EXISTS (SELECT 1 FROM product WHERE name = 'mobile');",
		},
		{
			name: "Combination check insert",
			entity: Entity{
				Table:   "product",
				Check:   []string{"name", "brand"},
				Columns: []string{"name", "brand", "description"},
				Values:  [][]any{{"mobile", "apple", "IPhone 16"}},
			},
			expected: "INSERT INTO product(name, brand, description) SELECT 'mobile', 'apple', 'IPhone 16' " +
				"WHERE NOT EXISTS (SELECT 1 FROM product WHERE name = 'mobile' AND brand = 'apple');",
		},
		{
			name: "Any value insert",
			entity: Entity{
				Table:   "product",
				Check:   []string{"name"},
				Columns: []string{"name", "price", "is_active"},
				Values:  [][]any{{"mobile", 10.5, true}},
			},
			expected: "INSERT INTO product(name, price, is_active) SELECT 'mobile', 10.5, true " +
				"WHERE NOT EXISTS (SELECT 1 FROM product WHERE name = 'mobile');",
		},
	}

	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			script := item.entity.GetScript()
			assert.Equal(t, item.expected, script)
		})
	}

}
