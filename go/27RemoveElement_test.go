package main

import (
	"math"
	"slices"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		expected []int
	}{
		{[]int{1, 2, 5, 3, 2}, 2, []int{1, 3, 5, math.MaxInt32, math.MaxInt32}},
		{[]int{3, 2, 2, 3}, 3, []int{2, 2, math.MaxInt32, math.MaxInt32}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 0, 1, 3, 4, math.MaxInt32, math.MaxInt32, math.MaxInt32}},
	}

	for _, tt := range tests {
		clone := slices.Clone(tt.a)
		res := removeElement(tt.a, tt.b)
		if !slices.Equal(tt.a, tt.expected) && len(tt.expected) != res {
			t.Errorf("removeElement(%v, %v) = %v; but got %v", clone, tt.b, tt.expected, tt.a)
		}
	}
}
