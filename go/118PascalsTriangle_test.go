package main

import (
	"testing"
)

func TestPascalsTriangle(t *testing.T) {
	tests := []struct {
		a        int
		expected [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}

	for _, tt := range tests {
		result := pascalsTriangle(tt.a)
		if !slicesEquals(tt.expected, result) {
			t.Errorf("pascalsTriangle(%v) = %v; but got %v", tt.a, tt.expected, result)
		}
	}
}
