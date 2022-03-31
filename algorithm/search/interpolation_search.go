package search

import (
	"github.com/anchorportal/go/algorithm/sort"
)

// 内插搜索，改良版的二分搜索，这个算法在数组中的值都是均匀分布时性能最好
func InterpolationSearch(arr []int, target int) int {
	arr = sort.QuickSort(arr)

	var low, high = 0, len(arr) - 1
	var delta = 0.0
	var position = 0

	for low <= high && target >= arr[low] && target <= arr[high] {
		delta = float64(target-arr[low]) / float64(arr[high]-arr[low])
		position = int(float64(low) + float64(high-low)*delta)

		if target < arr[position] {
			high = position - 1
		} else if target > arr[position] {
			low = position + 1
		} else {
			return position
		}
	}

	return -1
}
