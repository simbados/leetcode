package main

import (
	"testing"
)

func TestFindPeak(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
	}

	for _, tt := range tests {
		res := findPeakElement(tt.a)
		if res != tt.expected {
			t.Errorf("merge(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
