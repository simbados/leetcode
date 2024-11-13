package main

import (
	"slices"
	"testing"
)

func TestPascalsTriangle2(t *testing.T) {
	tests := []struct {
		a        int
		expected []int
	}{
		{4, []int{1, 4, 6, 4, 1}},
		{0, []int{1}},
	}

	for _, tt := range tests {
		result := pascalsTriangle2(tt.a)
		if !slices.Equal(result, tt.expected) {
			t.Errorf("pascalsTriangle2(%v) = %v; but got %v", tt.a, tt.expected, result)
		}
	}
}
