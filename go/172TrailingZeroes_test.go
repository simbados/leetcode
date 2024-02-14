package main

import (
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		a        int
		expected int
	}{
		{3, 0},
		{5, 1},
		{0, 0},
		{30, 7},
		{625, 156},
	}

	for _, tt := range tests {
		res := trailingZeroes(tt.a)
		if res != tt.expected {
			t.Errorf("trailingZeroes(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
