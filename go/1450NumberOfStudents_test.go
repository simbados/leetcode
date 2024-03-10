package main

import "testing"

func TestNumberOfStudentsDoingHomeWork(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		c        int
		expected int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 7}, 4, 1},
		{[]int{4}, []int{4}, 4, 1},
	}

	for _, tt := range tests {
		res := busyStudent(tt.a, tt.b, tt.c)
		if res != tt.expected {
			t.Errorf("busyStudent(%v, %v, %v) = %v; but got %v", tt.a, tt.b, tt.c, tt.expected, res)
		}
	}
}
