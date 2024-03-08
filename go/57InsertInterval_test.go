package main

import (
	"slices"
	"testing"
)

func TestOverlappingInterval(t *testing.T) {
	tests := []struct {
		a        [][]int
		b        []int
		expected [][]int
	}{
    {[][]int{{1,3},{6,9}}, []int{2,5}, [][]int{{1,5},{6,9}}},
    {[][]int{{1,2},{3,6},{6,7},{8,10},{12,16}}, []int{4,8}, [][]int{{1,2},{3,10},{12,16}}},
	}

	for _, tt := range tests {
		res := insertInterval(tt.a, tt.b)
		if len(tt.expected) == len(res) {
			for index, val := range res {
				if !slices.Equal(val, tt.expected[index]) {
					t.Errorf("insertInterval(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
				}
			}
		} else {
        t.Errorf("insertInterval(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
		}
  }
}
