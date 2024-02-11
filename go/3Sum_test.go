package main

import (
	"slices"
	"testing"
)

func Test3Sum(t *testing.T) {
	tests := []struct {
		a        []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{1, -1, -1, 0}, [][]int{{-1, 0, 1}}},
	}

	for _, tt := range tests {
		res := threeSum(tt.a)
		if len(tt.expected) != len(res) {
			t.Errorf("threeSum(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
		for i, val := range tt.expected {
			if !slices.Equal(val, res[i]) {
				t.Errorf("threesum(%v) = %v; but got %v", tt.a, tt.expected, res)
			}
		}
	}
}
