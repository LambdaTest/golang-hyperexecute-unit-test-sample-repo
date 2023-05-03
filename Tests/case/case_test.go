package main

import (
	"strings"
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

func checkCase(s string) string {
	if s == strings.ToUpper(s) {
		return "uppercase"
	} else if s == strings.ToLower(s) {
		return "lowercase"
	} else if strings.ContainsAny(s, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") && strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyz") {
		return "camelcase"
	} else {
		return "none"
	}
}
