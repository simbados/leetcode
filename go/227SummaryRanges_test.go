package main

import (
	"slices"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		a        []int
		expected []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{}, []string{}},
	}

	for _, tt := range tests {
		res := summaryRanges(tt.a)
		if !slices.Equal(res, tt.expected) {
			t.Errorf("summaryRanges(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
