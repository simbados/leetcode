package main

import (
	"slices"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		a        []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}

	for _, tt := range tests {
		res := permute(tt.a)
		if len(tt.expected) != len(res) {
			t.Errorf("permutations(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
		for i, val := range tt.expected {
			if !slices.Equal(val, res[i]) {
				t.Errorf("permutations(%v) = %v; but got %v", tt.a, tt.expected, res)
			}
		}
	}
}
