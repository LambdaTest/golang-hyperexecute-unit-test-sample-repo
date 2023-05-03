package main

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	// Define a struct to represent each test case
	type testCase struct {
		radius float64
		expectedArea float64
	}
	
	// Define the test cases
	testCases := []testCase{
		{0, 0},
		{1, math.Pi},
		{2, 4 * math.Pi},
	}
	
	// Run each test case
	for _, tc := range testCases {
		area := math.Pi * tc.radius * tc.radius
		if area != tc.expectedArea {
			t.Errorf("Expected area of circle with radius %v to be %v, but got %v", tc.radius, tc.expectedArea, area)
		}
	}
}

func TestMain(t *testing.T) {
	t.Run("TestCircleArea", TestCircleArea)
}
