package main

import (
	"testing"
)

func TestGastStation(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected int
	}{
		// {[]int{1,2,3,4,5}, []int{3,4,5,1,2}, 3},
		// {[]int{2,3,4}, []int{3,4,3}, -1},
		// {[]int{4,5,2,6,5,3}, []int{3,2,7,3,2,9}, -1},
		{[]int{5,8,2,8}, []int{6,5,6,6}, 3},
	}

	for _, tt := range tests {
		res := canCompleteCircuit(tt.a, tt.b)
		if res != tt.expected {
			t.Errorf("canComplete(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
	}
}
