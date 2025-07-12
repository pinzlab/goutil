package track

import (
	"reflect"
	"time"
)

// ToUpdate creates a map of updated fields from the provided data struct.
// It captures non-nil pointer fields and includes the updater's information.
//
// Parameters:
// - data: A pointer to the struct containing fields to check for updates.
// - updatedBy: An optional pointer to a string indicating the user who updated the record.
//
// Returns:
// A map[string]interface{} with updated field names and values, along with
// "UpdatedBy" and "UpdatedAt" if `updatedBy` is provided.
//
// Example:
// var userUpdates User
// updates := gorm.ToUpdate(&userUpdates, updatedBy)
func ToUpdate(data interface{}, updatedBy *int64) map[string]interface{} {
	// Ensure data is a pointer
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Ptr {
		panic("data must be a pointer to a struct")
	}

	updates := make(map[string]interface{})

	updates["UpdatedAt"] = time.Now()
	if updatedBy != nil {
		updates["UpdatedBy"] = updatedBy
	}

	dataValue = dataValue.Elem() // Dereference the pointer

	// Iterate over the fields of the struct
	for i := 0; i < dataValue.NumField(); i++ {
		field := dataValue.Type().Field(i)
		value := dataValue.Field(i)

		// Check if the field is a pointer and is not nil
		if field.Type.Kind() == reflect.Ptr {
			if !value.IsNil() {
				updates[field.Name] = value.Interface()
			}
		} else if field.Name != "ID" {
			updates[field.Name] = value
		}
	}

	return updates
}
