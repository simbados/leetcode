package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{1, 2}, 3, []int{2, 1}},
	}

	for _, tt := range tests {
		clone := slices.Clone(tt.a)
		rotate(tt.a, tt.b)
		fmt.Println(tt.a, tt.b, clone, tt.expected)
		if !slices.Equal(tt.a, tt.expected) {
			t.Errorf("rotate(%v, %v) = %v; but got %v", clone, tt.b, tt.expected, tt.a)
		}
	}
}
