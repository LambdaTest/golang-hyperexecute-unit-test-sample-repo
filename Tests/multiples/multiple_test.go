package multiples

import (
	"testing"
)

func TestIsMultipleOfFive(t *testing.T) {
	// Test cases
	tests := []struct {
		name   string
		input  []int
		output bool
	}{
		{"empty array", []int{}, true},
		{"all multiples of five", []int{5, 10, 15, 20}, true},
		{"not all multiples of five", []int{2, 5, 7, 10}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call function under test
			result := isMultipleOfFive(tt.input)

			// Check if output is as expected
			if result != tt.output {
				t.Errorf("Unexpected result. Expected %v, but got %v", tt.output, result)
			}
		})
	}
}
