package main

import (
	"slices"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		a        []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-1, 2, 1, -4}, -4, -4},
		{[]int{0, 0, 0}, 1, 0},
		{[]int{0, 3, 97, 102, 200}, 300, 300},
	}

	for _, tt := range tests {
		before := slices.Clone(tt.a)
		res := threeSumClosest(tt.a, tt.target)
		if res != tt.expected {
			t.Errorf("threeSumClosest(%v, %v) = %v; but got %v", before, tt.target, tt.expected, res)
		}
	}
}
