package main

import (
	"testing"
)

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	for _, tt := range tests {
		res := longestConsecutive(tt.a)
		if res != tt.expected {
			t.Errorf("longestConsecutive(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
