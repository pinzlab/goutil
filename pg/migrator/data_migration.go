package migrator

import (
	"gorm.io/gorm"
)

// DataMigration defines a single database DataMigration, including metadata and
// a collection of DataMigration components such as enums, entities, foreign keys,
// and data inserts.
//
// Each DataMigration can specify dependencies in the form of raw SQL strings that
// must be executed first, followed by other components like enums, entities,
// unique indexes, foreign keys, stored procedures, and data insertions.
//
// Once all operations are successful, the DataMigration is recorded in a tracking table.
type DataMigration struct {
	Code        string    // Unique code identifier for the DataMigration (max 20 chars)
	Name        string    // Human-readable name for the DataMigration (max 100 chars)
	Description string    // Optional detailed description of the DataMigration (max 255 chars
	Data        []*Entity // Initial data to seed conditionally
}

// GetCode returns the unique identifier for the migration
func (m *DataMigration) GetCode() string {
	return m.Code
}

// GetName returns the human-readable name for the migration
func (m *DataMigration) GetName() string {
	return m.Name
}

// Execute performs the data migration within a transaction
func (m *DataMigration) Execute(tx *gorm.DB) error {
	// Insert default data
	for _, data := range m.Data {
		if err := tx.Exec(data.GetScript()).Error; err != nil {
			return err
		}
	}
	return nil
}
