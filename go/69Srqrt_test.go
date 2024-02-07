package main

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	tests := []struct {
		a        int
		expected int
	}{
		{4, 2},
		{8, 2},
	}

	for _, tt := range tests {
		res := binarySqrt(tt.a)
		if res != tt.expected {
			t.Errorf("sqrt(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
