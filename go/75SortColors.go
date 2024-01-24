package main

// Bubble Sort
func bubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				buffer := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = buffer
			}
		}
	}
}

// Quicksort
func quickSort(nums []int) {
	quickSortRec(nums, 0, len(nums)-1)
}

func quickSortRec(nums []int, low int, high int) {
	if low >= 0 && high >= 0 && high > low {
		pivot := partition(nums, low, high)
		quickSortRec(nums, 0, pivot)
		quickSortRec(nums, pivot+1, high)
	}
}

func partition(nums []int, low int, high int) int {
	pivot := nums[(low+high)/2]
	i := low
	j := high
	for i <= j {
		for nums[i] < pivot {
			i = i + 1
		}
		for nums[j] > pivot {
			j = j - 1
		}
		if i >= j {
			return j
		}
		buffer := nums[i]
		nums[i] = nums[j]
		nums[j] = buffer
		i++
		j--
	}
	return j
}
