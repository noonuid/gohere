package problem0063

// 63. 不同路径 II

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。

// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

// 网格中的障碍物和空位置分别用 1 和 0 来表示。

// 动态规划，优化空间复杂度。
func uniquePathsWithObstacles_optimization(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for row := 0; row < m; row++ {
		if obstacleGrid[row][0] == 1 {
			f[0] = 0
		}
		for col := 1; col < n; col++ {
			if obstacleGrid[row][col] == 1 {
				f[col] = 0
			} else {
				f[col] = f[col] + f[col-1]
			}
		}
	}
	return f[n-1]
}

// 动态规划。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([][]int, m)
	f[0] = make([]int, n)
	if obstacleGrid[0][0] == 0 {
		f[0][0] = 1
	}
	for row := 1; row < m; row++ {
		f[row] = make([]int, n)
		if obstacleGrid[row][0] != 1 {
			f[row][0] = f[row-1][0]
		}
	}
	for col := 1; col < n; col++ {
		if obstacleGrid[0][col] != 1 {
			f[0][col] = f[0][col-1]
		}
	}

	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if obstacleGrid[row][col] != 1 {
				f[row][col] = f[row-1][col] + f[row][col-1]
			}
		}
	}
	return f[m-1][n-1]
}
