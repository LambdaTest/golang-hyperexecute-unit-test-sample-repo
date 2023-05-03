package main

import (
	"testing"
)

func TestSum(t *testing.T) {
    tests := []struct {
        name     string
        numbers  []int
        expected int
    }{
        {"Positive numbers", []int{1, 2, 3, 4, 5}, 15},
        {"Negative numbers", []int{-1, -2, -3, -4, -5}, -15},
        {"Mixed numbers", []int{-1, 2, -3, 4, -5}, -3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := sum(tt.numbers)

            if result != tt.expected {
                t.Errorf("sum(%v) = %d; expected %d", tt.numbers, result, tt.expected)
            }
        })
    }
}

func sum(numbers []int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
