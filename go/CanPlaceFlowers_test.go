package main

import (
	"testing"
)

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		expected bool
	}{
		{[]int{1,0,0,0,1}, 1, true},
		{[]int{1,0,0,0,1}, 2, false},
		{[]int{1,0,1,0,1,0,0}, 1, true},
		{[]int{0,1,0,0,0,0,0}, 2, true},
		{[]int{0,1,0,0,0,0,0}, 3, false},
	}

	for _, tt := range tests {
		res := canPlaceFlowers(tt.a, tt.b)
        if res != tt.expected {
            t.Errorf("canPlaceFlowers(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
        }
	}
}

