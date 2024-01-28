package main

import (
	"encoding/json"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	a := ListNode{Val: 3}
	b := ListNode{Val: 4, Next: &a}
	c := ListNode{Val: 2, Next: &b}
	d := ListNode{Val: 4}
	e := ListNode{Val: 6, Next: &d}
	f := ListNode{Val: 5, Next: &e}
	expectedA := ListNode{Val: 8}
	expectedB := ListNode{Val: 0, Next: &expectedA}
	expectedC := ListNode{Val: 7, Next: &expectedB}
	res := addTwoNumbers(&c, &f)
	if !listEquals(res, &expectedC) {
		marshalRes, _ := json.Marshal(*res)
		marshalExp, _ := json.Marshal(expectedC)
		t.Errorf("addTwoNumbers(%v, %v) = %v; but got %v", c, f, string(marshalExp), string(marshalRes))
	}
}

func listEquals(a *ListNode, b *ListNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if a.Val == b.Val {
		return listEquals(a.Next, b.Next)
	}
	return false
}
