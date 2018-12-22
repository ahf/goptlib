package pt

import (
	"bytes"
	"fmt"
	"sort"
	"unicode"
)

func kvline_value_needs_escape(input string) bool {
	for _, c := range input {
		if c == '\'' || c == '"' || !unicode.IsPrint(c) || unicode.IsSpace(c) {
			return true
		}
	}

	return false
}

func kvline_escape_value(input string) string {
	if !kvline_value_needs_escape(input) {
		return input
	}

	// This code could benefit from the strings.Builder, but we cannot use that
	// because we want to work on Debian releases that have older Go versions.
	var result bytes.Buffer

	// Worst case all characters needs to be escaped.
	result.Grow(2 * len(input))

	for _, c := range input {
		switch c {
		case '\'':
			result.WriteRune('\\')
			result.WriteRune('\'')
		case '"':
			result.WriteRune('\\')
			result.WriteRune('"')
		case '\n':
			result.WriteRune('\\')
			result.WriteRune('n')
		case '\t':
			result.WriteRune('\\')
			result.WriteRune('t')
		case '\r':
			result.WriteRune('\\')
			result.WriteRune('r')
		default:
			result.WriteRune(c)
		}
	}

	return fmt.Sprintf("\"%s\"", result.String())
}

func kvline_encode(input map[string]string) string {
	// We need to make sure our ordering is stable. Go's map's are for sensible
	// reasons non-stable.
	keys := make([]string, 0, len(input))

	for key := range input {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// This code could be refactored into using Go 1.10+ `strings.Builder`
	// type, but for now we can perfectly fine with the additional allocated
	// done by doing `+=` on strings :-)
	var result string

	for index := range keys {
		key := keys[index]
		value := input[key]

		// Are we not the first iteration we need to add a space.
		if index != 0 {
			result += " "
		}

		result += fmt.Sprintf("%s=%s",
			key, kvline_escape_value(value))
	}

	return result
}
