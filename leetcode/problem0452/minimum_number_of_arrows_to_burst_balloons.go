package problem0452

import (
	"cmp"
	"slices"
)

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	result := 0
	i, n := 0, len(points)
	for i < n {
		right := points[i][1]
		j := i + 1
		for j < n && points[j][0] <= right {
			right = min(right, points[j][1])
			j++
		}
		result++
		i = j
	}
	return result
}

func findMinArrowShots_greedy(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
	result := 1
	maxRight := points[0][1]
	for _, point := range points {
		if point[0] > maxRight {
			result++
			maxRight = point[1]
		}
	}
	return result
}
