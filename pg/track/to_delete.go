package track

import (
	"time"
)

// ToSoftDelete generates a map of fields used to perform a soft delete operation.
//
// It sets the "DeletedAt" field to the current time, indicating when the deletion occurred.
// If a non-nil user ID is provided via the `deletedBy` parameter, it also sets the "DeletedBy" field.
//
// This map can be used with ORM update functions (e.g., GORM) to mark a record as soft-deleted.
//
// Parameters:
//   - deletedBy: A pointer to an int representing the ID of the user performing the delete. Can be nil.
//
// Returns:
//   - A map[string]interface{} containing the fields to update for a soft delete.
func ToSoftDelete(deletedBy *int64) map[string]interface{} {
	updates := make(map[string]interface{})

	updates["DeletedAt"] = time.Now()
	if deletedBy != nil {
		updates["DeletedBy"] = deletedBy
	}

	return updates
}
