package main

import (
	"testing"
)

func TestReversePolish(t *testing.T) {
	tests := []struct {
		a        []string
		expected int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for _, tt := range tests {
		res := evalRPN(tt.a)
		if tt.expected != res {
			t.Errorf("evalRPN(%v) = %v; but got %v", tt.a, tt.expected, res)
		}
	}
}
