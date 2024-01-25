package main

import (
	"slices"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		a        [][]int
		expected [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {0, 5}}, [][]int{{0, 5}}},
	}

	for _, tt := range tests {
		res := mergeIntervals(tt.a)
		if len(tt.expected) == len(res) {
			for index, val := range res {
				if !slices.Equal(val, tt.expected[index]) {
					t.Errorf("merge(%v) = %v; but got %v", tt.a, tt.expected, res)
				}
			}
		} else {
			t.Errorf("merge(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
