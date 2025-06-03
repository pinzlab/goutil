package migrator

import (
	"gorm.io/gorm"
)

// Migration defines a single database migration, including metadata and
// a collection of migration components such as enums, entities, foreign keys,
// and data inserts.
//
// Each migration can specify dependencies in the form of raw SQL strings that
// must be executed first, followed by other components like enums, entities,
// unique indexes, foreign keys, stored procedures, and data insertions.
//
// Once all operations are successful, the migration is recorded in a tracking table.
type Migration struct {
	Code         string        // Unique code identifier for the migration
	Name         string        // Human-readable name for the migration
	Dependencies []string      // Raw SQL dependencies to execute before other steps
	Enums        []*Enum       // ENUM types to be created conditionally
	Entities     []interface{} // GORM models to be auto-migrated
	Uniques      []*Unique     // Unique constraints to be added via raw SQL
	ForeignKeys  []*Foreign    // Foreign key constraints to be added via raw SQL
	Procedures   []string      // Stored procedures or functions in SQL
	Data         []*Entity     // Initial data to seed conditionally
}

// Transaction applies the migration within the provided GORM database transaction.
//
// It performs all migration steps in a defined order:
// 1. Execute dependencies (raw SQL).
// 2. Create ENUM types if they do not exist.
// 3. Auto-migrate GORM entity models.
// 4. Add unique constraints using raw SQL.
// 5. Add foreign key constraints.
// 6. Execute stored procedures or SQL scripts.
// 7. Insert initial data conditionally.
//
// If any step fails, the transaction is rolled back and the error is returned.
// On success, the migration is recorded in a tracking table.
func (m *Migration) Transaction(tx *gorm.DB) error {

	// Migrating DEPENDENCIES for the database.
	for _, dep := range m.Dependencies {
		if err := tx.Exec(dep).Error; err != nil {
			return err
		}
	}

	// Migrating ENUM types to the database.
	for _, enum := range m.Enums {
		if err := tx.Exec(enum.GetScript()).Error; err != nil {
			return err
		}
	}

	// Migrating ENTITIES to the database.
	if len(m.Entities) > 0 {
		if err := tx.AutoMigrate(m.Entities...); err != nil {
			return err
		}
	}

	// Migrating UNIQUE constraints to the database.
	for _, unique := range m.Uniques {
		if err := tx.Exec(unique.GetScript()).Error; err != nil {
			return err
		}
	}

	// Migrating FOREIGN KEY constraints to the database.
	for _, fk := range m.ForeignKeys {
		if err := tx.Exec(fk.GetScript()).Error; err != nil {
			return err
		}
	}

	// Migrating STORED PROCEDURES to the database.
	for _, procedure := range m.Procedures {
		if err := tx.Exec(procedure).Error; err != nil {
			return err
		}
	}

	// Migrating INITIAL DATA to the database.
	for _, data := range m.Data {
		if err := tx.Exec(data.GetScript()).Error; err != nil {
			return err
		}
	}

	// Save the migration record to the tracking table.
	if err := tx.Create(&tracker{Code: m.Code, Name: m.Name}).Error; err != nil {
		return err
	}

	return nil
}
