package main

import (
	"testing"
)

func TestNumberOfBits(t *testing.T) {
	tests := []struct {
		a        uint32
		expected int
	}{
		{0b00000000000000000000000000001011, 3},
		{0b00000000000000000000000010000000, 1},
		{0b11111111111111111111111111111101, 31},
	}

	for _, tt := range tests {
		res := hammingWeight(tt.a)
		if res != tt.expected {
			t.Errorf("longestConsecutive(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
