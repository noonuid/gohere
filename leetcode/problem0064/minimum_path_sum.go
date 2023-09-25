package problem0064

import "math"

// 64. 最小路径和

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 说明：每次只能向下或者向右移动一步。

// 动态规划，并使用滚动数组优化空间复杂度。
func minPathSum_dp_scrolling_array(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// f 为滚动数组。
	f := make([]int, n)
	f[0] = grid[0][0]
	for col := 1; col < n; col++ {
		f[col] = f[col-1] + grid[0][col]
	}
	for row := 1; row < m; row++ {
		f[0] = f[0] + grid[row][0]
		for col := 1; col < n; col++ {
			if f[col] < f[col-1] {
				f[col] = f[col] + grid[row][col]
			} else {
				f[col] = f[col-1] + grid[row][col]
			}
		}
	}
	return f[n-1]
}

// 动态规划。
func minPathSum_dynamic_programming(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	f[0][0] = grid[0][0]
	for row := 1; row < m; row++ {
		f[row][0] = f[row-1][0] + grid[row][0]
	}
	for col := 1; col < n; col++ {
		f[0][col] = f[0][col-1] + grid[0][col]
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if f[row-1][col] < f[row][col-1] {
				f[row][col] = f[row-1][col] + grid[row][col]
			} else {
				f[row][col] = f[row][col-1] + grid[row][col]
			}
		}
	}
	return f[m-1][n-1]
}

// 回溯法。
// 超出时间限制。
func minPathSum_backtrack(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := math.MaxInt

	var backtrack func(row, col, pathSum int)
	backtrack = func(row, col, pathSum int) {
		if row == m || col == n {
			return
		}
		if sum := pathSum + grid[row][col]; sum < res {
			if row == m-1 && col == n-1 {
				res = sum
			}
			backtrack(row+1, col, sum)
			backtrack(row, col+1, sum)
		}
	}

	backtrack(0, 0, 0)
	return res
}
