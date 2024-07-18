package main

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World!", "!dlroW olleH"},
		{"kayak", "kayak"},
		{"qwerty", "ytrewq"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"racecar", "racecar"},
		{"12345", "54321"},
	}

	for _, test := range tests {
		result := reverseString(test.input)
		if result != test.expected {
			t.Errorf("reverse_string(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
