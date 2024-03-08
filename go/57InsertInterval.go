package main

import "fmt"

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	copy := append([][]int{newInterval}, intervals...)
  fmt.Println(copy)
	for i := 1; i < len(copy); i++ {
		currentVal := copy[i]
		previousVal := copy[i-1]
		if (previousVal[0] >= currentVal[0] && previousVal[0] <= currentVal[1]) ||
			(previousVal[1] >= currentVal[0] && previousVal[1] <= currentVal[1]) || 
      (previousVal[0] <= currentVal[0] && previousVal[1] >= currentVal[1]) {
			newValue := []int{min(previousVal[0], currentVal[0]), max(previousVal[1], currentVal[1])}
			copy[i] = newValue
			copy = removeIndex(i-1, copy)
      i--
		} else if (currentVal[1] < previousVal[1]) {
      cache := copy[i-1]
      copy[i-1] = copy[i]
      copy[i] = cache
    }
	}
	return copy
}

func removeIndex(index int, array [][]int) [][]int {
	arrayClone := make([][]int, 0, len(array)-1)
	arrayClone = append(arrayClone, array[:index]...)
	arrayClone = append(arrayClone, array[index+1:]...)
	return arrayClone
}
