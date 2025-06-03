package migrator

import (
	"github.com/pinzlab/goutil/terminal"
	"gorm.io/gorm"
)

// migrator manages the execution of a sequence of database migrations
// in a PostgreSQL database using GORM. It ensures that each migration
// is applied exactly once by recording them in a tracking table.
type migrator struct {
	db     *gorm.DB     // Database connection
	schema []*Migration // List of migrations to apply
}

// New creates a new migrator instance with the given database connection
// and a variadic list of migration definitions.
//
// Parameters:
//   - db: a pointer to the gorm.DB instance (PostgreSQL GORM wrapper)
//   - items: variadic list of pointers to Migration structs
//
// Returns:
//   - *migrator: a configured migrator ready to run migrations
func New(db *gorm.DB, items ...*Migration) *migrator {
	return &migrator{
		db:     db,
		schema: items,
	}
}

// trackerExists checks whether the internal 'migrations' tracking table
// exists in the connected PostgreSQL database.
//
// Returns:
//   - bool: true if the table exists, false otherwise
//   - error: if the query fails
func (m *migrator) trackerExists() (bool, error) {
	exists := false
	tableName := (&tracker{}).TableName()

	query := `
        SELECT EXISTS (
            SELECT FROM information_schema.tables 
            WHERE table_schema = 'public' AND table_name = ?
        );
    `
	err := m.db.Raw(query, tableName).Scan(&exists).Error
	return exists, err
}

// migrateTracker ensures that the 'migrations' tracking table exists.
// If it does not, the function creates it using GORM's AutoMigrate.
//
// Returns:
//   - error: if the migration operation fails
func (m *migrator) migrateTracker() error {
	te, err := m.trackerExists()
	if err != nil {
		return err
	}

	if !te {
		terminal.Info("Migrate Tracker")
		if err := m.db.AutoMigrate(&tracker{}); err != nil {
			return err
		}
	}
	return nil
}

// checkMigration determines whether a migration with the given code
// has already been applied, by checking the tracking table.
//
// Parameters:
//   - code: the unique migration code to check
//
// Returns:
//   - bool: true if the migration has already been run
//   - error: if the database query fails
func (m *migrator) checkMigration(code string) (bool, error) {
	var count int64
	err := m.db.Model(&tracker{}).Where("code = ?", code).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Run applies all pending migrations in the order they were defined.
// It first ensures the tracker table exists, then checks each migration
// by code. If a migration has not been applied, it is executed inside
// a database transaction. Successfully applied migrations are recorded.
//
// Returns:
//   - error: if any migration step fails
func (m *migrator) Run() error {
	err := m.migrateTracker()
	if err != nil {
		return err
	}

	for _, migration := range m.schema {
		exists, err := m.checkMigration(migration.Code)
		if err != nil {
			return err
		}

		if !exists {
			terminal.About("Migrate", migration.Name)
			err := m.db.Transaction(migration.Transaction)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
