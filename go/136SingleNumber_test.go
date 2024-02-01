package main

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, tt := range tests {
		res := singleNumber(tt.a)
		if res != tt.expected {
			t.Errorf("singleNumber(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
