package main

type ListNode struct {
	Val  int       `json:"Val"`
	Next *ListNode `json:"Next"`
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	over := 0
	node := ListNode{}
	nextNode := &node
	for l1 != nil || l2 != nil || over != 0 {
		if l1 == nil && l2 == nil {
			nextNode.Val, over = addValue(0, 0, over)
		} else if l1 == nil {
			nextNode.Val, over = addValue(0, l2.Val, over)
			l2 = l2.Next
		} else if l2 == nil {
			nextNode.Val, over = addValue(l1.Val, 0, over)
			l1 = l1.Next
		} else {
			nextNode.Val, over = addValue(l1.Val, l2.Val, over)
			l1 = l1.Next
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil || over != 0 {
			temp := ListNode{}
			nextNode.Next = &temp
			nextNode = &temp
		}
	}
	return &node
}

func addValue(a int, b int, over int) (int, int) {
	currentAdd := (a + b) + over
	val := currentAdd % 10
	if currentAdd > 9 {
		over = 1
	} else {
		over = 0
	}
	return val, over
}
