package main

import (
    "testing"
)

func TestIntegerFunctions(t *testing.T) {
    t.Run("Test integer length", func(t *testing.T) {
        input := 12345
        expectedLength := 5

        result := lenInt(input)

        if result != expectedLength {
            t.Errorf("Expected length of %d but got %d", expectedLength, result)
        }
    })

    t.Run("Test integer addition", func(t *testing.T) {
        input1 := 10
        input2 := 20
        expectedSum := 30

        result := addInt(input1, input2)

        if result != expectedSum {
            t.Errorf("Expected sum of %d but got %d", expectedSum, result)
        }
    })
}

func lenInt(n int) int {
    // function to calculate the length of an integer
    if n == 0 {
        return 1
    }

    count := 0
    for n != 0 {
        n /= 10
        count++
    }

    return count
}

func addInt(a, b int) int {
    // function to add two integers
    return a + b
}
