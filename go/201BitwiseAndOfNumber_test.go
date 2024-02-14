package main

import (
	"testing"
)

func TestBitwiseAndOfNumber(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{5, 7, 4},
		{0, 0, 0},
		{1, 2147483647, 0},
		{0, 1, 0},
		{6, 7, 6},
		{12, 14, 12},
	}

	for _, tt := range tests {
		res := rangeBitwiseAnd(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("rangeBitwiseAnd(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
