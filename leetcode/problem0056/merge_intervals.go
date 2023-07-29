package problem0056

import (
	"sort"
)

// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

// 将区间先按左端点排序，再依次合并。
func merge(intervals [][]int) [][]int {
	// 将区间按左端点排序。
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	merged := [][]int{intervals[0]}
	for _, interval := range intervals {
		if interval[0] > merged[len(merged)-1][1] {
			merged = append(merged, interval)
		} else {
			// 合并区间。
			if interval[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = interval[1]
			}
		}
	}
	return merged
}
