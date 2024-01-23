package main

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, tt := range tests {
		res := searchInsert(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("searchInsert(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
