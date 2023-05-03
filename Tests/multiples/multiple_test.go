package main

import (
    "testing"
)

func TestIsMultipleOf5(t *testing.T) {
    testCases := []struct {
        name     string
        numbers  []int
        expected bool
    }{
        {"empty array", []int{}, true},
        {"all multiples of 5", []int{0, 5, 10, 15}, true},
        {"mixed numbers", []int{1, 3, 5, 7}, false},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            actual := true
            for _, num := range tc.numbers {
                if num%5 != 0 {
                    actual = false
                    break
                }
            }
            if actual != tc.expected {
                t.Errorf("expected %v, but got %v", tc.expected, actual)
            }
        })
    }
}
