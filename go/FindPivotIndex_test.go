package main

import (
	"testing"
)

func TestFindPivotIndex(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}

	for _, tt := range tests {
		res := pivotIndex(tt.a)
		if res != tt.expected {
			t.Errorf("pivotIndex(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
