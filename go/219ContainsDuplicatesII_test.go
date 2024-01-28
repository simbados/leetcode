package main

import (
	"testing"
)

func TestContainsDuplicatesII(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, tt := range tests {
		res := containsNearbyDuplicate(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("containsNearbyDuplicate(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
