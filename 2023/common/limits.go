package common

import "math"

func MaxInts(arr []int) int {
	max := math.MinInt

	for _, e := range arr {
		if e > max {
			max = e
		}
	}

	return max
}

func MaxIntInts(matrix [][]int) int {
	max := math.MinInt

	for _, arr := range matrix {
		for _, e := range arr {
			if e > max {
				max = e
			}
		}
	}

	return max
}
