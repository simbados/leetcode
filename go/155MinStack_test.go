package main

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	a := Constructor2()
	a.Push(1)
	a.Push(2)

	res := a.GetMin()
	if res != 1 {
		t.Errorf("MinStack GetMin() = %v; but got %v", 1, res)
	}
	a.Pop()
	res = a.Top()
	if res != 1 {
		t.Errorf("MinStack Top() = %v; but got %v", 1, res)
	}
	a.Pop()
	if a.Len() != 0 {
		t.Errorf("MinStack Pop = %v; but got %v", 2, res)
	}
}
