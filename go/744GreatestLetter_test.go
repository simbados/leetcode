package main

import (
	"testing"
)

func TestGreatestLetter(t *testing.T) {
	tests := []struct {
		a        []byte
		b        byte
		expected byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
	}

	for _, val := range tests {
		res := nextGreatestLetter(val.a, val.b)
		if res != val.expected {
			t.Errorf("longestPalindrome(%v, %v) = %v; but got %v", val.a, val.b, val.expected, res)
		}
	}
}
