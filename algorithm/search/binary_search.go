package search

import "sort"

// 二分搜索
func BinarySearch(arr []int, target int) int {
	sort.Ints(arr)

	var low, high = 0, len(arr) - 1

	for low <= high {
		var middle = (low + high) / 2
		if target < arr[middle] {
			high = middle - 1
		} else if target > arr[middle] {
			low = middle + 1
		} else {
			return middle
		}
	}

	return -1
}
