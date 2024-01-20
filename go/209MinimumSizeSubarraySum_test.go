package main

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		a        int
		b        []int
		expected int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}, 8},
	}

	for _, tt := range tests {
		result := minSubArrayLen(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("minSubArrayLen(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, result)
		}
	}
}
