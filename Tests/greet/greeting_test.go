package main

import (
    "fmt"
    "testing"
)

func TestGreeting(t *testing.T) {
    fmt.Println("Greetings!")
}

func TestGreetingWithSubtests(t *testing.T) {
    t.Run("Test Greeting", func(t *testing.T) {
        fmt.Println("Greetings!")
    })
}
