package main

import (
	"slices"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
    boolinger := make([]bool, len(candies))
    if len(candies) == 0 {
        return boolinger
    }
    sorted := slices.Clone(candies)
    slices.Sort(sorted)
    highest := sorted[len(sorted) - 1]
    for index, value := range(candies) {
        boolinger[index] = value + extraCandies >= highest
    }
    return boolinger
}
