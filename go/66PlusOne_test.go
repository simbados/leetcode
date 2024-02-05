package main

import (
	"slices"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		a        []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, tt := range tests {
		res := plusOne(tt.a)
		if !slices.Equal(res, tt.expected) {
			t.Errorf("plusOne(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
