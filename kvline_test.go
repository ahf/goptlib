package pt

import (
	"testing"
)

func TestKVLineEncoder(t *testing.T) {
	tests := [...]struct {
		value    map[string]string
		expected string
	}{
		{map[string]string{"A": "B", "CD": "EF"}, "A=B CD=EF"},
		{map[string]string{"AB": "C", "CDE": "F G"}, "AB=C CDE=\"F G\""},
		{map[string]string{"A": "Foo Bar Baz\r\t\n\"'"}, "A=\"Foo Bar Baz\\r\\t\\n\\\"\\'\""},
	}

	for _, input := range tests {
		encoded := kvline_encode(input.value)
		if input.expected != encoded {
			t.Errorf("kvline_encode(%v) → \"%v\" (expected \"%v\")",
				input.value, encoded, input.expected)
		}
	}
}
