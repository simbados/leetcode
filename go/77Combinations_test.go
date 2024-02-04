package main

import (
	"slices"
	"testing"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected [][]int
	}{
		{5, 4, [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}}},
	}

	for _, tt := range tests {
		res := combine(tt.a, tt.b)
		if !slicesEquals(res, tt.expected) {
			t.Errorf("combine(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}

func slicesEquals(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if !slices.Equal(val, b[i]) {
			return false
		}
	}
	return true
}
