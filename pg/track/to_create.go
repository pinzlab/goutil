package track

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// injectCreatedFields sets CreatedAt (and CreatedBy if available) on structs embedding Create or CreateOnly.
// It uses reflection to find and populate these metadata fields during creation.
func injectCreatedFields(entity interface{}, createdBy *int64) {
	targetValue := reflect.ValueOf(entity).Elem()

	for i := 0; i < targetValue.NumField(); i++ {
		field := targetValue.Field(i)
		fieldType := targetValue.Type().Field(i)

		if field.Kind() != reflect.Struct {
			continue
		}

		switch fieldType.Type {
		case reflect.TypeOf(Create{}):
			if createdBy != nil {
				createdAt := field.FieldByName("CreatedAt")
				if createdAt.IsValid() && createdAt.CanSet() {
					createdAt.Set(reflect.ValueOf(time.Now()))
				}

				createdByField := field.FieldByName("CreatedBy")
				if createdByField.IsValid() && createdByField.CanSet() {
					createdByField.SetInt(int64(*createdBy))
				}
			}

			return

		case reflect.TypeOf(CreateOnly{}):
			createdAt := field.FieldByName("CreatedAt")
			if createdAt.IsValid() && createdAt.CanSet() {
				createdAt.Set(reflect.ValueOf(time.Now()))
			}

			return
		}
	}
}

// convertValue handles the conversion of values between different types (e.g., string -> int, string -> float, etc.)
// It attempts to convert the source field value into the destination field value type.
func convertValue(sourceFieldVal reflect.Value, destFieldVal reflect.Value) {
	if sourceFieldVal.Kind() == reflect.String {
		// Convert string to appropriate destination field type
		switch destFieldVal.Kind() {
		case reflect.Int64:
			// Convert string to int64
			intVal, err := strconv.ParseInt(sourceFieldVal.String(), 10, 64)
			if err == nil {
				destFieldVal.SetInt(intVal)
			} else {
				fmt.Println("Error converting string to int64:", err)
			}

		case reflect.Int32:
			// Convert string to int32
			intVal, err := strconv.ParseInt(sourceFieldVal.String(), 10, 32)
			if err == nil {
				destFieldVal.SetInt(intVal)
			} else {
				fmt.Println("Error converting string to int32:", err)
			}

		case reflect.Float64:
			// Convert string to float64
			floatVal, err := strconv.ParseFloat(sourceFieldVal.String(), 64)
			if err == nil {
				destFieldVal.SetFloat(floatVal)
			} else {
				fmt.Println("Error converting string to float64:", err)
			}

		case reflect.Bool:
			// Convert string to bool
			boolVal, err := strconv.ParseBool(sourceFieldVal.String())
			if err == nil {
				destFieldVal.SetBool(boolVal)
			} else {
				fmt.Println("Error converting string to bool:", err)
			}

		default:
			// Handle custom types (e.g., ProfileType)
			if destFieldVal.Type().Kind() == reflect.String {
				destFieldVal.SetString(sourceFieldVal.String())
			}
		}
	} else {
		// Direct assignment for non-string types
		destFieldVal.Set(sourceFieldVal)
	}
}

// ToCreate copies data from the input struct to the target entity struct, preparing it for database creation.
//
// It matches fields by name and handles ID fields ending with "ID" by converting string IDs to int32 if necessary.
// Additionally, it sets creation metadata fields `CreatedAt` and `CreatedBy` if the target entity embeds
// the `Create` or `CreateOnly` structs.
//
// Parameters:
//   - input: Pointer to the source struct containing data to copy.
//   - entity: Pointer to the destination struct where data will be copied.
//   - createdBy: Optional pointer to an int representing the creator's identifier.
//     If provided and the entity has a `Create` struct, `CreatedBy` will be set.
//
// Notes:
//   - Both input and entity must be pointers to structs.
//   - If `createdBy` is nil, only `CreatedAt` will be set if the entity has a `CreateOnly` struct.
//   - Fields that are pointers in the source will be dereferenced or copied appropriately.
func ToCreate(input, entity interface{}, createdBy *int64) {
	// Obtain reflect.Values for source and target
	sourceValue := reflect.ValueOf(input)
	targetValue := reflect.ValueOf(entity)

	// Check if both source and target are pointers to structs
	if sourceValue.Kind() != reflect.Ptr || targetValue.Kind() != reflect.Ptr {
		fmt.Println("Both source and destination must be pointers to structs")
		return
	}

	// Dereference the pointers to obtain the underlying structs
	sourceElem := sourceValue.Elem()
	targetElem := targetValue.Elem()

	// Get the types of the source and target structs
	sourceType := sourceElem.Type()
	targetType := targetElem.Type()

	// Iterate through fields of the source struct
	for i := 0; i < sourceElem.NumField(); i++ {
		sourceField := sourceType.Field(i)

		// Assign ID source to ID target
		if strings.HasSuffix(sourceField.Name, "ID") {
			for j := 0; j < targetElem.NumField(); j++ {
				targetField := targetType.Field(j)
				if targetField.Name == sourceField.Name {
					if sourceField.Type.Kind() == reflect.String && targetField.Type.Kind() == reflect.Int64 {
						// Convert string to int64
						sourceFieldVal := sourceElem.Field(i)
						destFieldVal := targetElem.Field(j)
						intVal, err := strconv.ParseInt(sourceFieldVal.Interface().(string), 10, 64) // Cambiar a int64
						if err != nil {
							fmt.Println("Error converting string to int64:", err)
							continue
						}
						destFieldVal.SetInt(intVal)
					}
					break

				}
			}
			continue
		}

		// Assign another field from source to target
		for j := 0; j < targetElem.NumField(); j++ {
			targetField := targetType.Field(j)
			if sourceField.Name == targetField.Name {
				sourceFieldVal := sourceElem.Field(i)
				destFieldVal := targetElem.Field(j)
				// Handle pointer types (copy or dereference)
				if sourceFieldVal.Kind() == reflect.Ptr && !sourceFieldVal.IsNil() {
					if destFieldVal.Kind() == reflect.Ptr {
						destFieldVal.Set(sourceFieldVal)
					} else {
						destFieldVal.Set(sourceFieldVal.Elem())
					}
				} else if sourceFieldVal.Kind() != reflect.Ptr {
					// Handle non-pointer types with direct assignment
					convertValue(sourceFieldVal, destFieldVal)
				}
				break
			}
		}
	}

	// Set CreatedAt and CreatedBy if present in the entity
	injectCreatedFields(entity, createdBy)
}
