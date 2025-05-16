package track

import (
	"time"
)

// Delete represents metadata for tracking deletion operations.
type Delete struct {
	// DeletedAt is the timestamp when the record was deleted (soft delete).
	DeletedAt *time.Time `gorm:"column:dat;type:timestamptz;null"`

	// DeletedBy is the identifier of the user who deleted the record.
	DeletedBy *int `gorm:"column:cby;type:integer;null"`
}
