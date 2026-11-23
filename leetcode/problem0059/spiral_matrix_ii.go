package problem0059

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col, dirIdx := 0, 0, 0
	for num := 1; num <= n*n; num++ {
		matrix[row][col] = num
		dir := dirs[dirIdx]
		if r, c := row+dir[0], col+dir[1]; !(-1 < r && r < n) || !(-1 < c && c < n) || matrix[r][c] > 0 {
			dirIdx = (dirIdx + 1) % 4
			dir = dirs[dirIdx]
		}
		row, col = row+dir[0], col+dir[1]
	}
	return matrix
}

func generateMatrix_level(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	top, left, bottom, right := 0, 0, n-1, n-1
	num := 1
	for top < bottom && left < right {
		for col := left; col < right; col++ {
			matrix[top][col] = num
			num++
		}
		for row := top; row < bottom; row++ {
			matrix[row][right] = num
			num++
		}
		for col := right; col > left; col-- {
			matrix[bottom][col] = num
			num++
		}
		for row := bottom; row > top; row-- {
			matrix[row][left] = num
			num++
		}
		top, left = top+1, left+1
		bottom, right = bottom-1, right-1
	}
	if n%2 == 1 {
		matrix[top][left] = num
	}
	return matrix
}
