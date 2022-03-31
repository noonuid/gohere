package problem0074

import "sort"

// 整体二分查找
func searchMatrix(matrix [][]int, target int) bool {
	var m, n = len(matrix), len(matrix[0])
	var low, high = 0, m*n - 1
	var middle = (low + high) / 2

	for low <= high {
		var row, col = middle / n, middle % n
		if target < matrix[row][col] {
			high = middle - 1
		} else if target > matrix[row][col] {
			low = middle + 1
		} else {
			return true
		}
		middle = (low + high) / 2
	}
	return false
}

// 两次二分查找
func searchMatrix1(matrix [][]int, target int) bool {
	var row = sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	var col = sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
