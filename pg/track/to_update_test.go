package track

import (
	"testing"

	"github.com/pinzlab/goutil/internal/helper"
	"github.com/stretchr/testify/assert"
)

type UpdateData struct {
	Name      *string
	Age       *int
	Email     *string
	Address   *string
	ExtraInfo *string
}

func TestToUpdate(t *testing.T) {
	tests := []struct {
		name         string
		data         UpdateData
		updatedBy    *int64
		expectedKeys []string
	}{
		{
			name:         "Only non-nil pointer fields",
			data:         UpdateData{Name: helper.Pointer("Alice"), Age: helper.Pointer(30), ExtraInfo: helper.Pointer("Extra")},
			updatedBy:    helper.Pointer[int64](1),
			expectedKeys: []string{"UpdatedBy", "UpdatedAt", "Name", "Age", "ExtraInfo"},
		},
		{
			name:         "No updatedBy and some nil fields",
			data:         UpdateData{Email: helper.Pointer("bob@example.com")},
			updatedBy:    nil,
			expectedKeys: []string{"Email"},
		},
		{
			name:         "Only updatedBy with nil data fields",
			data:         UpdateData{},
			updatedBy:    helper.Pointer[int64](1),
			expectedKeys: []string{"UpdatedBy", "UpdatedAt"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			updates := ToUpdate(&test.data, test.updatedBy)

			// Check if the expected keys are present in the updates map
			for _, key := range test.expectedKeys {
				assert.Contains(t, updates, key)
			}

			// Additionally check if UpdatedAt is set if updatedBy is provided
			if test.updatedBy != nil {
				assert.Contains(t, updates, "UpdatedAt")
			}
		})
	}
}
