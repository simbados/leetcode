package main

import (
	"testing"
)

func TestBuySell2(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		res := maxProfit2(tt.a)
		if res != tt.expected {
			t.Errorf("maxProfit2(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
