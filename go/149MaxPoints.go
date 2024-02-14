package main

import (
	"math"
)

func maxPoints(points [][]int) int {
	if len(points) == 2 {
		return 2
	}
	res := 0
	for index := 0; index < len(points)-2; index++ {
		cache := 0
		mapping := map[float64]int{}
		for j := index + 1; j <= len(points)-1; j++ {
			gradient := calcGradient(points[index], points[j])
			val, exists := mapping[gradient]
			if !exists {
				mapping[gradient] = 1
			} else {
				mapping[gradient] = val + 1
			}
			if mapping[gradient] > cache {
				cache = mapping[gradient]
			}
		}
		res = max(res, cache)
	}
	return res + 1
}

func calcGradient(a []int, b []int) float64 {
	x := a[0] - b[0]
	y := a[1] - b[1]
	if x == 0 {
		return 0
	}
	if y == 0 {
		return math.Inf(1)
	}
	return float64(y) / float64(x)
}
