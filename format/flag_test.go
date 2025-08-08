package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlag(t *testing.T) {
	// Test cases
	tests := []struct {
		input    string
		expected string
	}{
		// Basic cases
		{"Hello World", "hello-world"},
		{"Go Language", "go-language"},

		// Accented characters
		{"Café", "cafe"},
		{"Résumé", "resume"},
		{"Llamé", "llame"},
		{"Über", "uber"},

		// Mixed cases
		{"  Multiple   Spaces   ", "multiple-spaces"},
		{"Mixed CASES And Áccents", "mixed-cases-and-accents"},

		// Edge cases
		{"NoAccentsOrSpaces", "noaccentsorspaces"},
	}

	for _, item := range tests {
		t.Run(item.input, func(t *testing.T) {
			flag := Flag(item.input)
			assert.Equal(t, item.expected, flag, "Expected %q but got %q", item.expected, flag)
		})
	}
}
