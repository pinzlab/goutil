package pg

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetUniqueScript tests the GetScript method of the Unique struct.
func TestGetUniqueScript(t *testing.T) {
	// Define test cases for different scenarios
	tests := []struct {
		name     string // Name of the test case
		unique   Unique // The Unique object to test
		expected string // The expected SQL script output
	}{
		{
			name: "Simple unique index",
			unique: Unique{
				Table:   "users",
				Columns: []string{"email"},
			},
			expected: `
	CREATE UNIQUE INDEX IF NOT EXISTS uni_users_email
		ON public.users(email)
	WHERE dat IS NULL;
	`,
		},
		{
			name: "Multiple columns in unique index",
			unique: Unique{
				Table:   "orders",
				Columns: []string{"user_id", "order_date"},
			},
			expected: `
	CREATE UNIQUE INDEX IF NOT EXISTS uni_orders_user_id_order_date
		ON public.orders(user_id, order_date)
	WHERE dat IS NULL;
	`,
		},
		{
			name: "Multiple columns with underscores in index name",
			unique: Unique{
				Table:   "products",
				Columns: []string{"category", "sku"},
			},
			expected: `
	CREATE UNIQUE INDEX IF NOT EXISTS uni_products_category_sku
		ON public.products(category, sku)
	WHERE dat IS NULL;
	`,
		},
	}

	// Run each test case
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			// Generate the script and compare with expected output
			script := item.unique.GetScript()
			assert.Equal(t, strings.TrimSpace(item.expected), strings.TrimSpace(script))
		})
	}
}
