package repeat

import (
    "testing"
)

func TestRepeatedA(t *testing.T) {
    testCases := []struct {
        word string
        expectRepeated bool
    }{
        {"aaaaaa", true},
        {"aa", true},
        {"aaa", true},
    }

    for _, tc := range testCases {
        tc := tc // capture range variable
        t.Run(tc.word, func(t *testing.T) {
            if hasRepeatedA(tc.word) != tc.expectRepeated {
                t.Errorf("Expected word '%s' to have repeated 'a'=%v", tc.word, tc.expectRepeated)
            }
        })
    }
}

