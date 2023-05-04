package hello

import "testing"

func TestReverse(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"hello", "Hello", "olleH"},
        {"empty", "", ""},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := reverse(tc.input)
            if result != tc.expected {
                t.Errorf("reverse(%s) = %s; want %s", tc.input, result, tc.expected)
            }
        })
    }
}

