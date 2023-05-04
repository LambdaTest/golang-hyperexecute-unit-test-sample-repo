package palindrome

import (
	"fmt"
    "testing"
)

func TestIsPalindrome(t *testing.T) {
    cases := []struct {
        input    string
        expected bool
    }{
        {"racecar", true},
        {"hello", false},
        {"", true},
    }

    for _, c := range cases {
        t.Run(fmt.Sprintf("Input: %s", c.input), func(t *testing.T) {
            result := isPalindrome(c.input)
            if result != c.expected {
                t.Errorf("Expected isPalindrome(%q) to be %v, but got %v", c.input, c.expected, result)
            }
        })
    }
}


