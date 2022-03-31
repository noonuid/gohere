package problem0048

// 方法一：利用辅助二维数组
func rotate(matrix [][]int) {
	var m = [][]int{}
	for i := 0; i < len(matrix); i++ {
		m = append(m, matrix[i])
	}
	for i := 0; i < len(m); i++ {
		var row = []int{}
		for j := len(m[i]) - 1; j >= 0; j-- {
			row = append(row, m[j][i])
		}
		matrix[i] = row
	}
}

// 方法二：先水平翻转，后沿对角线翻转
func rotateFlip(matrix [][]int) {
	var n = len(matrix)

	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}

	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
