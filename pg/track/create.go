package track

import (
	"time"
)

// Create represents metadata for tracking creation operations.
type Create struct {
	// CreatedAt is the timestamp when the record was created.
	CreatedAt time.Time `gorm:"column:cat;type:timestamptz;default:now();not null"`

	// CreatedBy is the identifier of the creator.
	CreatedBy int64 `gorm:"column:cby;type:integer;not null"`
}

// CreateOnly represents metadata for tracking creation timestamp only.
type CreateOnly struct {
	// CreatedAt is the timestamp when the record was created.
	CreatedAt time.Time `gorm:"column:cat;type:timestamptz;default:now();not null"`
}
