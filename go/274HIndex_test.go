package main

import (
	"testing"
)

func TestHIndex(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
	}

	for _, tt := range tests {
		res := hIndex(tt.a)
		if res != tt.expected {
			t.Errorf("hIndex(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
