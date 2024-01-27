package main

import (
	"testing"
)

func TestBuySell(t *testing.T) {
	tests := []struct {
		a        []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		res := maxProfit(tt.a)
		if res != tt.expected {
			t.Errorf("maxProfit(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
