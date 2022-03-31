package problem0240

import "sort"

// 暴力破解
func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// 每一行使用一次二分法
func searchMatrixBinarySearch(matrix [][]int, target int) bool {
	for _, row := range matrix {
		var index = sort.SearchInts(row, target)
		if index < len(row) && row[index] == target {
			return true
		}
	}
	return false
}

// Z 字形查找
func searchMatrixZ(matrix [][]int, target int) bool {
	var m, n = len(matrix), len(matrix[0])
	var x, y = 0, n - 1
	for x <= m-1 && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else if matrix[x][y] < target {
			x++
		}
	}
	return false
}
