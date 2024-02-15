package main

import (
	"testing"
)

func TestHouseRobber(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
	}

	for _, tt := range tests {
		res := rob(tt.a)
		if res != tt.expected {
			t.Errorf("rob(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
