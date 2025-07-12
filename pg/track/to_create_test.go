package track

import (
	"testing"

	"github.com/pinzlab/goutil/internal/helper"
	"github.com/stretchr/testify/assert"
)

type ClientType string
type GQLClientType string

const (
	ClientCompany    ClientType    = "Company"
	GQLClientCompany GQLClientType = "Company"
)

type NewClient struct {
	Type     GQLClientType
	Name     string
	EstabID  string
	PersonID *string
	Address  *string
}

type Client struct {
	Type     ClientType
	ID       int64
	EstabID  int64
	PersonID int64
	Name     string
	Address  *string
	Create   Create
}

func TestToCreate(t *testing.T) {
	tests := []struct {
		name      string
		input     NewClient
		createdBy *int64
		expected  Client
	}{
		{
			name:      "Successful mapping with CreatedBy and Address",
			input:     NewClient{Name: "Alice", EstabID: "5", PersonID: helper.Pointer("456"), Address: helper.Pointer("123 Main St")},
			createdBy: helper.Pointer[int64](1),
			expected:  Client{Name: "Alice", EstabID: 5, PersonID: 457, Address: helper.Pointer("123 Main St"), Create: Create{CreatedBy: 1}},
		},
		{
			name:     "Successful mapping without CreatedBy and Address",
			input:    NewClient{Name: "Bob", PersonID: helper.Pointer("101")},
			expected: Client{Name: "Bob", PersonID: 101, Address: nil, Create: Create{}},
		},
		{
			name:     "Successful mapping enum",
			input:    NewClient{Type: GQLClientCompany, PersonID: helper.Pointer("101")},
			expected: Client{Type: ClientCompany, PersonID: 101, Create: Create{}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var entity Client

			ToCreate(&test.input, &entity, test.createdBy)

			// Assert the expected values
			assert.Equal(t, test.expected.Name, entity.Name)
			if test.createdBy != nil {
				assert.NotZero(t, entity.Create.CreatedAt)
				assert.Equal(t, *test.createdBy, entity.Create.CreatedBy)
			} else {
				assert.Zero(t, entity.Create.CreatedAt)
				assert.Empty(t, entity.Create.CreatedBy)
			}

			// Assert Address field
			if test.input.Address != nil {
				assert.Equal(t, *test.input.Address, *entity.Address)
			} else {
				assert.Nil(t, entity.Address)
			}
		})
	}
}
