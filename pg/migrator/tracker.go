package migrator

import "time"

// tracker represents a record in the 'migrations' table used to track the
// execution of applied migrations. This helps prevent reapplying the same
// migration multiple times.
type tracker struct {
	// Code is a unique identifier for the migration.
	// It is the primary key and is limited to 20 characters.
	Code string `gorm:"primaryKey;type:varchar(20)"`

	// CreatedAt stores the timestamp of when the migration was applied.
	// It is mapped to the 'cat' column in the database and defaults to the current time.
	CreatedAt time.Time `gorm:"column:cat;type:timestamptz;default:now();not null"`

	// Name is a human-readable name or label for the migration.
	// It is required and limited to 100 characters.
	Name string `gorm:"type:varchar(100);not null"`

	// Description is an field that provides more details about the migration.
	// It can be null and is limited to 255 characters.
	Description string `gorm:"type:varchar(255);not null"`
}

// TableName overrides the default GORM table name for the tracker struct.
// It specifies that migration tracking records are stored in the "migrations" table.
func (*tracker) TableName() string {
	return "migrations"
}
