package pg

import (
	"regexp"
	"strings"

	"gorm.io/gorm/clause"
)

// Enum to specify sorting order.
type Order string

const (
	// Ascending order (e.g., from low to high).
	OrderAsc Order = "Asc"
	// Descending order (e.g., from high to low).
	OrderDesc Order = "Desc"
)

// OrderBy creates a GORM `OrderByColumn` clause for SQL `ORDER BY` statements. It either
// converts a CamelCase column name to snake_case or uses a mapped column name from the
// provided `definition` list.
//
// Parameters:
//   - order: Sorting direction (`OrderAsc` or `OrderDesc`).
//   - column: Column name in CamelCase format.
//   - definition: Optional column name mappings in the form of "original.mapped".
//
// Returns:
//   - A `clause.OrderByColumn` with the mapped or converted column name and sorting direction.
//
// Example:
//
//	result := r.OrderBy(OrderDesc, "username")  // Converts to "user_name"
//	definition := []string{"user.user_name"}
//	result := r.OrderBy(OrderAsc, "username", definition)  // Uses "user_name"
func NewOrder(orderInput interface{}, column string, definition ...string) clause.OrderByColumn {
	var order Order

	// Determine order type
	switch v := orderInput.(type) {
	case Order:
		order = v
	case string:
		order = Order(v)
	default:
		// Fallback to Asc if unknown type
		order = OrderAsc
	}

	if len(definition) > 0 {
		for _, itemDefinition := range definition {
			parts := strings.Split(itemDefinition, ".")

			if parts[1] == column {
				return clause.OrderByColumn{
					Desc:   order == OrderDesc,
					Column: clause.Column{Name: itemDefinition},
				}
			}
		}

	}

	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(column, "${1}_${2}")

	return clause.OrderByColumn{
		Desc:   order == OrderDesc,
		Column: clause.Column{Name: strings.ToLower(snake)},
	}
}
