package main

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tt := range tests {
		res := twoSum(tt.a, tt.b)
		if !slices.Equal(res, tt.expected) {
			t.Errorf("twoSum(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
