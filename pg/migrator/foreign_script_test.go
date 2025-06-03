package migrator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetForeignScript tests the GetScript method of the Foreign type.
func TestGetForeignScript(t *testing.T) {
	tests := []struct {
		name     string  // name of the test case
		foreign  Foreign // the Foreign object to test
		expected string  // the expected SQL script
	}{
		{
			name: "Simple foreign key",
			foreign: Foreign{
				Table:       "orders",
				ForeignID:   "customer_id",
				Reference:   "customers",
				ReferenceID: "id",
			},
			expected: `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname= 'fk_orders_customer_id') THEN
			ALTER TABLE public.orders ADD CONSTRAINT fk_orders_customer_id
			FOREIGN KEY (customer_id) REFERENCES customers(id) 
			ON UPDATE CASCADE ON DELETE CASCADE;
		END IF;
	END $$;
	`,
		},
		{
			name: "Foreign key with different table",
			foreign: Foreign{
				Table:       "products",
				ForeignID:   "category_id",
				Reference:   "categories",
				ReferenceID: "id",
			},
			expected: `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname= 'fk_products_category_id') THEN
			ALTER TABLE public.products ADD CONSTRAINT fk_products_category_id
			FOREIGN KEY (category_id) REFERENCES categories(id) 
			ON UPDATE CASCADE ON DELETE CASCADE;
		END IF;
	END $$;
	`,
		},
		{
			name: "Foreign key with single column",
			foreign: Foreign{
				Table:       "employees",
				ForeignID:   "department_id",
				Reference:   "departments",
				ReferenceID: "id",
			},
			expected: `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname= 'fk_employees_department_id') THEN
			ALTER TABLE public.employees ADD CONSTRAINT fk_employees_department_id
			FOREIGN KEY (department_id) REFERENCES departments(id) 
			ON UPDATE CASCADE ON DELETE CASCADE;
		END IF;
	END $$;
	`,
		},
	}

	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			script := item.foreign.GetScript()
			assert.Equal(t, item.expected, script)
		})
	}
}
