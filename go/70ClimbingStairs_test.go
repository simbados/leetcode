package main

import (
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		a        int
		expected int
	}{
		{2, 2},
		{3, 3},
		{5, 8},
	}

	for _, tt := range tests {
		res := climbStairs(tt.a)
		if res != tt.expected {
			t.Errorf("climbingStairs(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
