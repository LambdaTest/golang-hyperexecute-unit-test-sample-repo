package main

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

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

