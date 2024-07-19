package main

import (
	"slices"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		a        []int
		b        int
		expected []bool
	}{
		{[]int{2,3,5,1,3}, 3, []bool{true,true,true,false,true}},
		{[]int{4,2,1,1,2}, 1, []bool{true,false,false,false,false}},
	}

	for _, tt := range tests {
		res := kidsWithCandies(tt.a, tt.b)
        if !slices.Equal(res, tt.expected) {
            t.Errorf("kidsWithCandies(%v, %v) = %v; but got %v", tt.a, tt.b, tt.expected, res)
        }
	}
}

