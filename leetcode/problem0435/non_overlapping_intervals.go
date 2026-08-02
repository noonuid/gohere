package problem0435

import (
	"slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}
	dpMax := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if intervals[j][1] <= intervals[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dpMax < dp[i] {
			dpMax = dp[i]
		}
	}
	return n - dpMax
}

func eraseOverlapIntervals_greedy(intervals [][]int) int {
	n := len(intervals)
	slices.SortFunc(intervals, func(a, b []int) int { return a[1] - b[1] })
	preRight := intervals[0][1]
	count := 1
	for i := 1; i < n; i++ {
		if preRight <= intervals[i][0] {
			preRight = intervals[i][1]
			count++
		}
	}
	return n - count
}
