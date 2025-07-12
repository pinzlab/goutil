package track

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test when deletedBy is provided
func TestToSoftDelete_WithDeletedBy(t *testing.T) {
	userID := int64(42)
	result := ToSoftDelete(&userID)

	deletedAt, ok := result["DeletedAt"].(time.Time)
	assert.True(t, ok, "DeletedAt should be of type time.Time")
	assert.False(t, deletedAt.IsZero(), "DeletedAt should be set to current time")

	assert.Equal(t, &userID, result["DeletedBy"], "DeletedBy should be set to the provided user ID")
}

// Test when deletedBy is nil
func TestToSoftDelete_WithoutDeletedBy(t *testing.T) {
	result := ToSoftDelete(nil)

	deletedAt, ok := result["DeletedAt"].(time.Time)
	assert.True(t, ok, "DeletedAt should be of type time.Time")
	assert.False(t, deletedAt.IsZero(), "DeletedAt should be set to current time")

	_, exists := result["DeletedBy"]
	assert.False(t, exists, "DeletedBy should not be set when deletedBy is nil")
}
