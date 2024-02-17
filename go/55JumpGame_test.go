package main

import (
	"testing"
)

func TestJumpGame(t *testing.T) {
	tests := []struct {
		a        []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{2, 0}, true},
	}

	for _, tt := range tests {
		res := canJump(tt.a)
		if res != tt.expected {
			t.Errorf("canJump(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
