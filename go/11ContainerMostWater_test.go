package main

import (
	"testing"
)

func TestContainerMostWater(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 4, 3}, 4},
	}

	for _, tt := range tests {
		res := maxArea(tt.a)
		if res != tt.expected {
			t.Errorf("maxArea(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
