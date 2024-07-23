package main

import (
	"testing"
)

func TestFindTheHighestAltitude(t *testing.T) {
	tests := []struct {
		a        []int
		expected        int
	}{
		{[]int{-5,1,5,0,-7}, 1 },
	}

	for _, tt := range tests {
		res := largestAltitude(tt.a)
        if res != tt.expected {
            t.Errorf("largestAltitude(%v) = %v; but got %v", tt.a, tt.expected, res)
        }
	}
}

