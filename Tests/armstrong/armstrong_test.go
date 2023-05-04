package armstrong

import (
	"fmt"
    "testing"
)

func TestIsArmstrongNumber(t *testing.T) {
    testCases := []struct {
        input    int
        expected bool
    }{
        {0, true},
        {1, true},
        {153, true},
        {370, true},
        {371, true},
        {407, true},
        {1634, true},
        {8208, true},
        {9474, true},
        {10, false},
        {100, false},
        {121, false},
        {200, false},
        {400, false},
        {1000, false},
    }

    for _, tc := range testCases {
        t.Run(fmt.Sprintf("Input:%v,Expected:%v", tc.input, tc.expected), func(t *testing.T) {
            result := isArmstrongNumber(tc.input)
            if result != tc.expected {
                t.Errorf("isArmstrongNumber(%d) = %v; expected %v", tc.input, result, tc.expected)
            }
        })
    }
}
