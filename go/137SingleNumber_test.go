package main

import (
	"testing"
)

func TestSingleNumberII(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{2, 2, 3, 2}, 3},
		{[]int{0, 1, 0, 1, 0, 1, 99}, 99},
	}

	for _, tt := range tests {
		result := singleNumberII(tt.a)
		if result != tt.expected {
			t.Errorf("maxPoints(%v) = %v; but got %v", tt.a, tt.expected, result)
		}
	}
}
