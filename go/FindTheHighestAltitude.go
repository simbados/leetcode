package main

import (
	"slices"
)

func largestAltitude(gain []int) int {
	altitudeArr := []int{}
	altitudeArr = append(altitudeArr, 0)
	for i, v := range gain {
		altitudeArr = append(altitudeArr, v+altitudeArr[i])
	}
	slices.Sort(altitudeArr)
	return altitudeArr[len(altitudeArr)-1]
}
