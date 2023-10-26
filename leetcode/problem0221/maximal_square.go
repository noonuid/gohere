package problem0221

// 221. 最大正方形

// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

// 动态规划。
func maximalSquare_dynamic_programming(matrix [][]byte) int {
	maxEdge := 0
	m, n := len(matrix), len(matrix[0])
	f := make([][]int, m)
	for row := 0; row < m; row++ {
		f[row] = make([]int, n)
		f[row][0] = int(matrix[row][0] - '0')
		if f[row][0] > maxEdge {
			maxEdge = f[row][0]
		}
	}
	for col := 0; col < n; col++ {
		f[0][col] = int(matrix[0][col] - '0')
		if f[0][col] > maxEdge {
			maxEdge = f[0][col]
		}
	}
	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][col] == '1' {
				f[row][col] = min(min(f[row-1][col], f[row][col-1]), f[row-1][col-1]) + 1

				if f[row][col] > maxEdge {
					maxEdge = f[row][col]
				}
			}
		}
	}
	return maxEdge * maxEdge
}
