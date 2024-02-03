package main

func sortedArrayToBST(nums []int) *TreeNode {
	return addChilds(nums, 0, len(nums)-1)
}

func addChilds(array []int, start int, end int) *TreeNode {
	if end < start {
		return nil
	}
	middle := end - (end-start)/2
	root := TreeNode{Val: array[middle]}
	root.Left = addChilds(array, start, middle-1)
	root.Right = addChilds(array, middle+1, end)
	return &root
}
