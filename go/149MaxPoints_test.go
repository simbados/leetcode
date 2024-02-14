package main

import (
	"testing"
)

func TestMaxPoints(t *testing.T) {
	tests := []struct {
		a        [][]int
		expected int
	}{
		{[][]int{{1, 1}, {2, 2}, {3, 3}}, 3},
		{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
	}

	for _, tt := range tests {
		result := maxPoints(tt.a)
		if result != tt.expected {
			t.Errorf("maxPoints(%v) = %v; but got %v", tt.a, tt.expected, result)
		}
	}
}
