package _go

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m < m+n {
		nums1[m] = nums2[n-1]
		m++
		n--
	}
	sort.Ints(nums1)
}
