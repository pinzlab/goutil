package format

import (
	"strings"
)

// Slug generates a URL-friendly slug from the given input string.
// It replaces spaces with hyphens, removes accents, and converts
// the result to lowercase.
//
// Example:
//
//	input  := "Hola Señor Mundo"
//	output := Slug(input)
//	// output: "hola-senor-mundo"
func Slug(input string) string {
	// Replace spaces with hyphens
	result := strings.ReplaceAll(strings.ToLower(input), " ", "-")

	// Create a mapping of accented characters to their non-accented counterparts
	accentMap := map[rune]rune{
		'à': 'a', 'á': 'a', 'â': 'a', 'ã': 'a', 'ä': 'a', 'å': 'a',
		'ç': 'c',
		'è': 'e', 'é': 'e', 'ê': 'e', 'ë': 'e',
		'ì': 'i', 'í': 'i', 'î': 'i', 'ï': 'i',
		'ñ': 'n',
		'ò': 'o', 'ó': 'o', 'ô': 'o', 'õ': 'o', 'ö': 'o',
		'ù': 'u', 'ú': 'u', 'û': 'u', 'ü': 'u',
		'ý': 'y', 'ÿ': 'y',
	}

	// Create a function to replace accented characters with their non-accented counterparts
	replace := func(r rune) rune {
		if val, ok := accentMap[r]; ok {
			return val
		}
		return r
	}

	// Use strings.Map to replace accented characters in the string
	result = strings.Map(replace, result)

	// Normalize multiple hyphens to a single hyphen
	result = normalizeHyphens(result)

	return result
}

// normalizeHyphens replaces multiple consecutive hyphens with a single hyphen
func normalizeHyphens(s string) string {
	return strings.Join(strings.FieldsFunc(s, func(r rune) bool {
		return r == '-'
	}), "-")
}
