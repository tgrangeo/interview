package main

import "testing"

func TestPalindrome(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"kayak", true},
        {"qwerty", false},
        {"racecar", true},
        {"level", true},
        {"hello", false},
        {"", true},            // Empty string
        {"a", true},           // Single character
        {"ab", false},         // Two different characters
        {"aa", true},          // Two same characters
        {"abcba", true},       // Odd length palindrome
        {"abccba", true},      // Even length palindrome
    }

    for _, test := range tests {
        result := palindrome(test.input)
        if result != test.expected {
            t.Errorf("palindrome(%q) = %v; want %v", test.input, result, test.expected)
        }
    }
}
