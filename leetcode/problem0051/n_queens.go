package problem0051

// 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
func solveNQueens(n int) [][]string {
	result := [][]string{}
	path := make([][]byte, n)
	for row := 0; row < n; row++ {
		b := make([]byte, n)
		for col := 0; col < n; col++ {
			b[col] = '.'
		}
		path[row] = b
	}
	// columns，diagonals1，diagonals2 分别记录已放置皇后的列，“/”斜线，“\”斜线。
	columns, diagonals1, diagonals2 := map[int]struct{}{}, map[int]struct{}{}, map[int]struct{}{}
	backtrack(&result, path, 0, columns, diagonals1, diagonals2)
	return result
}

func backtrack(result *[][]string, path [][]byte, row int, columns, diagonals1, diagonals2 map[int]struct{}) {
	if row == len(path) {
		// 复制 path。
		dst := make([]string, len(path))
		for i := 0; i < len(path); i++ {
			dst[i] = string(path[i])
		}
		// 将新的解答添加到最终结果中。
		*result = append(*result, dst)
		return
	}

	for col := 0; col < len(path); col++ {
		// 检查同一列。
		if _, ok := columns[col]; ok {
			continue
		}
		// 检查同一“/”斜线。
		diagonal1 := row + col
		if _, ok := diagonals1[diagonal1]; ok {
			continue
		}
		// 检查同一“\”斜线。
		diagonal2 := row - col
		if _, ok := diagonals2[diagonal2]; ok {
			continue
		}
		// 在 col 处放置皇后并更新 columns，diagonals1，diagonals2。
		path[row][col] = 'Q'
		columns[col], diagonals1[diagonal1], diagonals2[diagonal2] = struct{}{}, struct{}{}, struct{}{}
		backtrack(result, path, row+1, columns, diagonals1, diagonals2)
		// 还原。
		path[row][col] = '.'
		delete(columns, col)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}
