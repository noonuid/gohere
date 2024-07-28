package problem0073

// 73. 矩阵置零

// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

// 使用矩阵的第一行、第一列和额外的标记变量存储含有 0 的行和列。
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for _, row := range matrix {
		if row[0] == 0 {
			col0 = true
			break
		}
	}
	for row := 0; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[row][col] == 0 {
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}
	for row := m - 1; row >= 0; row-- {
		for col := 1; col < n; col++ {
			if matrix[row][0] == 0 || matrix[0][col] == 0 {
				matrix[row][col] = 0
			}
		}
		if col0 {
			matrix[row][0] = 0
		}
	}
}

// 枚举。
func setZeroes_enum(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zeroRow, zeroCol := make([]bool, m), make([]bool, n)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				zeroRow[row] = true
				zeroCol[col] = true
			}
		}
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if zeroRow[row] || zeroCol[col] {
				matrix[row][col] = 0
			}
		}
	}
}
