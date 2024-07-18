package main

import "testing"

func TestReverseInteger(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{12345, 54321},
		{-12345, -54321},
		{120, 21},
		{0, 0},
		{1534236469, 0},   // Overflow case, should return 0
		{-1534236469, 0},  // Overflow case, should return 0
	}

	for _, test := range tests {
		result := reverseInteger(test.input)
		if result != test.expected {
			t.Errorf("reverse_integer(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
