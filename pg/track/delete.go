package track

import (
	"gorm.io/gorm"
)

// Delete represents metadata for tracking deletion operations.
type Delete struct {
	// DeletedAt is the timestamp when the record was deleted (soft delete).
	DeletedAt *gorm.DeletedAt `gorm:"column:dat;type:timestamptz;null"`

	// DeletedBy is the identifier of the user who deleted the record.
	DeletedBy *int64 `gorm:"column:cby;type:integer;null"`
}

// DeleteOnly represents metadata for soft deletion timestamp only.
type DeleteOnly struct {
	// DeletedAt is the timestamp when the record was deleted (soft delete).
	DeletedAt gorm.DeletedAt `gorm:"column:dat;type:timestamptz;null"`
}
