package track

import (
	"time"
)

// Update represents metadata for tracking update operations.
type Update struct {
	// UpdatedAt is the timestamp when the record was last updated.
	UpdatedAt *time.Time `gorm:"column:uat;type:timestamptz;null"`

	// UpdatedBy is the identifier of the user who last updated the record.
	UpdatedBy *int64 `gorm:"column:uby;type:integer;null"`
}

// UpdateOnly represents metadata for tracking update timestamp only.
type UpdateOnly struct {
	// UpdatedAt is the timestamp when the record was last updated.
	UpdatedAt *time.Time `gorm:"column:uat;type:timestamptz;null"`
}
