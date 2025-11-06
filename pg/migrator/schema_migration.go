package migrator

import (
	"gorm.io/gorm"
)

// SchemaMigration defines a single database SchemaMigration, including metadata and
// a collection of SchemaMigration components such as enums, entities, foreign keys,
// and data inserts.
//
// Each SchemaMigration can specify dependencies in the form of raw SQL strings that
// must be executed first, followed by other components like enums, entities,
// unique indexes, foreign keys, stored procedures, and data insertions.
//
// Once all operations are successful, the SchemaMigration is recorded in a tracking table.
type SchemaMigration struct {
	Code         string        // Unique code identifier for the SchemaMigration (max 20 chars)
	Name         string        // Human-readable name for the SchemaMigration (max 100 chars)
	Description  string        // Optional detailed description of the SchemaMigration (max 255 chars)
	Dependencies []string      // Raw SQL dependencies to execute before other steps
	Enums        []*Enum       // ENUM types to be created conditionally
	Entities     []interface{} // GORM models to be auto-migrated
	Uniques      []*Unique     // Unique constraints to be added via raw SQL
	ForeignKeys  []*Foreign    // Foreign key constraints to be added via raw SQL
	Procedures   []string      // Stored procedures or functions in SQL
}

// GetCode returns the unique identifier for the migration
func (m *SchemaMigration) GetCode() string {
	return m.Code
}

// GetName returns the human-readable name for the migration
func (m *SchemaMigration) GetName() string {
	return m.Name
}

// Execute performs the schema migration within a transaction
func (m *SchemaMigration) Execute(tx *gorm.DB) error {
	// Execute dependencies
	for _, dep := range m.Dependencies {
		if err := tx.Exec(dep).Error; err != nil {
			return err
		}
	}

	// Create ENUM types
	for _, enum := range m.Enums {
		if err := tx.Exec(enum.GetScript()).Error; err != nil {
			return err
		}
	}

	// Auto-migrate entities
	if len(m.Entities) > 0 {
		if err := tx.AutoMigrate(m.Entities...); err != nil {
			return err
		}
	}

	// Add unique constraints
	for _, unique := range m.Uniques {
		if err := tx.Exec(unique.GetScript()).Error; err != nil {
			return err
		}
	}

	// Add foreign key constraints
	for _, fk := range m.ForeignKeys {
		if err := tx.Exec(fk.GetScript()).Error; err != nil {
			return err
		}
	}

	// Execute stored procedures
	for _, procedure := range m.Procedures {
		if err := tx.Exec(procedure).Error; err != nil {
			return err
		}
	}

	// Save the migration record to the tracking table.
	if err := tx.Create(&tracker{Code: m.Code, Name: m.Name, Description: m.Description}).Error; err != nil {
		return err
	}

	return nil
}
