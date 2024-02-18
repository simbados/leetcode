package main

import (
	"testing"
)

func TestJumpGame2(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 1, 0, 4}, 2},
		{[]int{1, 3, 4, 3, 1, 2, 1, 2}, 4},
	}

	for _, tt := range tests {
		res := jump(tt.a)
		if res != tt.expected {
			t.Errorf("jump(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
