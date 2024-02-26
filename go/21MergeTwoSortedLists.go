package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	start := result
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val > list2.Val {
				result.Next = &ListNode{Val: list2.Val}
                list2 = list2.Next
			} else {
				result.Next = &ListNode{Val: list1.Val}
                list1 = list1.Next
			}
		} else if list1 == nil {
			result.Next = &ListNode{Val: list2.Val}
            list2 = list2.Next
		} else {
			result.Next = &ListNode{Val: list1.Val}
            list1 = list1.Next
        }
        result = result.Next
	}
	return start.Next
}
