package main

import (
	"slices"
	"testing"
)

func TestProductOfArray(t *testing.T) {
	tests := []struct {
		a        []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, tt := range tests {
		res := productExceptSelf(tt.a)
		if !slices.Equal(res, tt.expected) {
			t.Errorf("%v = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
