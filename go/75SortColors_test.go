package main

import (
	"slices"
	"testing"
)

func testData() []struct {
	a        []int
	expected []int
} {
	return []struct {
		a        []int
		expected []int
	}{
		{[]int{1, 2, 5, 3, 2}, []int{1, 2, 2, 3, 5}},
		{[]int{3, 2, 2, 3}, []int{2, 2, 3, 3}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, []int{0, 0, 1, 2, 2, 2, 3, 4}},
	}
}

func TestBubbleSort(t *testing.T) {
	tests := testData()
	for _, tt := range tests {
		clone := slices.Clone(tt.a)
		bubbleSort(tt.a)
		if !slices.Equal(tt.a, tt.expected) {
			t.Errorf("sortColors(%v) = %v; but got %v", clone, tt.expected, tt.a)
		}
	}
}
