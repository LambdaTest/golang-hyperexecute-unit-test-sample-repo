package main

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

func isArmstrongNumber(n int) bool {
    numDigits := len(fmt.Sprint(n))
    sum := 0
    for i := n; i > 0; i /= 10 {
        digit := i % 10
        sum += pow(digit, numDigits)
    }
    return sum == n
}

func pow(x, y int) int {
    if y == 0 {
        return 1
    }
    return x * pow(x, y-1)
}
