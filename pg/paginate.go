package pg

// Paginate calculates the offset for a pagination query based on the current page and the number of items per page (take).
//
// The function computes the offset by multiplying the page number (page - 1) with the number of items per page (take).
// This offset can then be used in a database query to fetch the correct slice of data.
//
// Parameters:
//   - page: The current page number, starting from 1.
//   - take: The number of items to return per page.
//
// Returns:
//   - The offset to be used in the query, which is calculated as (page - 1) * take.
//
// Example:
//
//	page := 2
//	take := 10
//	offset := r.Paginate(page, take)
//	// offset will be 10 (for the second page, with 10 items per page).
func Paginate(page, take int) int {
	return (page - 1) * take

}
