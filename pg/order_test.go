package pg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/clause"
)

func TestOrder(t *testing.T) {
	tests := []struct {
		order       Order
		column      string
		definitions []string
		expected    clause.OrderByColumn
	}{
		{
			order:  OrderAsc,
			column: "UserName",
			expected: clause.OrderByColumn{
				Desc:   false,
				Column: clause.Column{Name: "user_name"},
			},
		},
		{
			order:  OrderDesc,
			column: "OrderAmount",
			expected: clause.OrderByColumn{
				Desc:   true,
				Column: clause.Column{Name: "order_amount"},
			},
		},
		{
			order:  OrderAsc,
			column: "CreatedAt",
			expected: clause.OrderByColumn{
				Desc:   false,
				Column: clause.Column{Name: "created_at"},
			},
		},
		{
			order:  OrderDesc,
			column: "LastUpdated",
			expected: clause.OrderByColumn{
				Desc:   true,
				Column: clause.Column{Name: "last_updated"},
			},
		},
		{
			order:       OrderDesc,
			column:      "username",
			definitions: []string{"user.username", "user.name"},
			expected: clause.OrderByColumn{
				Desc:   true,
				Column: clause.Column{Name: "user.username"},
			},
		},
	}

	for _, item := range tests {

		t.Run(item.column, func(t *testing.T) {
			result := NewOrder(item.order, item.column, item.definitions...)
			assert.Equal(t, item.expected.Desc, result.Desc, "Desc mismatch")
			assert.Equal(t, item.expected.Column.Name, result.Column.Name, "Column name mismatch")
		})
	}
}
