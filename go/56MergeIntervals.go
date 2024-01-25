package main

import (
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0] || (intervals[a][0] == intervals[b][0] && intervals[a][1] < intervals[b][1])
	})
	answer := [][]int{intervals[0]}
	for _, val := range intervals {
		if answer[len(answer)-1][1] < val[0] {
			answer = append(answer, val)
		} else {
			answer[len(answer)-1][1] = max(answer[len(answer)-1][1], val[1])
		}
	}
	return answer
}
