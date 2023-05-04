package greet

import (
    "testing"
)

func TestExtractWord(t *testing.T) {
    // Define the input array of words
    words := []string{"hello", "world", "extract", "this", "word"}

    // Call the ExtractWord function and store the result
    result := ExtractWord(words)

    // Check if the result is correct
    if result != "extract" {
        t.Errorf("Expected 'extract', but got '%s'", result)
    }
}

func main(t *testing.T) {
    // Run the test case
    t.Run("ExtractWord", TestExtractWord)
}
