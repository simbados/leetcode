package main

import (
	"encoding/json"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	a := ListNode{Val: 4}
	b := ListNode{Val: 2, Next: &a}
	c := ListNode{Val: 1, Next: &b}
	d := ListNode{Val: 4}
	e := ListNode{Val: 3, Next: &d}
	f := ListNode{Val: 1, Next: &e}
	expectedA := ListNode{Val: 4}
    expectedB := ListNode{Val: 4, Next: &expectedA}
    expectedC := ListNode{Val: 3, Next: &expectedB}
    expectedD := ListNode{Val: 2, Next: &expectedC}
    expectedE := ListNode{Val: 1, Next: &expectedD}
    expectedF := ListNode{Val: 1, Next: &expectedE}
	res := mergeTwoLists(&c, &f)
	if !listEquals(res, &expectedF) {
		marshalRes, _ := json.Marshal(*res)
		marshalExp, _ := json.Marshal(expectedF)
		t.Errorf("mergeTwoSortedLists(%v, %v) = %v; but got %v", c, f, string(marshalExp), string(marshalRes))
	}
}

