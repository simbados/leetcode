package main

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		c        []int
		d        int
		expected []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{}, 0, []int{}, 0, []int{}},
	}

	for _, tt := range tests {
		clone := slices.Clone(tt.a)
		merge(clone, tt.b, tt.c, tt.d)
		if !slices.Equal(clone, tt.expected) {
			t.Errorf("merge(%v, %v, %v, %v) = %v; but got %v", tt.a, tt.b, tt.c, tt.d, tt.expected, clone)
		}
	}
}
