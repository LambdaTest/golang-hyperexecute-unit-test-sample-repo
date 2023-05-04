package case1

import (
	"testing"
)

func TestCheckCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"HELLO WORLD", "uppercase"},
		{"hello world", "lowercase"},
		{"HelloWorld", "camelcase"},
		{"hElLo WoRlD", "camelcase"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := checkCase(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
