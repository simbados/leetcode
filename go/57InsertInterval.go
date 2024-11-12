package main

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	copyInterval := append([][]int{newInterval}, intervals...)
	for i := 1; i < len(copyInterval); i++ {
		currentVal := copyInterval[i]
		previousVal := copyInterval[i-1]
		if (previousVal[0] >= currentVal[0] && previousVal[0] <= currentVal[1]) ||
			(previousVal[1] >= currentVal[0] && previousVal[1] <= currentVal[1]) ||
			(previousVal[0] <= currentVal[0] && previousVal[1] >= currentVal[1]) {
			newValue := []int{min(previousVal[0], currentVal[0]), max(previousVal[1], currentVal[1])}
			copyInterval[i] = newValue
			copyInterval = removeIndex(i-1, copyInterval)
			i--
		} else if currentVal[1] < previousVal[1] {
			cache := copyInterval[i-1]
			copyInterval[i-1] = copyInterval[i]
			copyInterval[i] = cache
		}
	}
	return copyInterval
}

func removeIndex(index int, array [][]int) [][]int {
	arrayClone := make([][]int, 0, len(array)-1)
	arrayClone = append(arrayClone, array[:index]...)
	arrayClone = append(arrayClone, array[index+1:]...)
	return arrayClone
}
