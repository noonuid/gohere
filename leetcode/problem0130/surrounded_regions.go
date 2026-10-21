package problem0130

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if !(-1 < row && row < m && -1 < col && col < n) || board[row][col] != 'O' {
			return
		}
		board[row][col] = 'I'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := 1; i < m-1; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'I':
				board[i][j] = 'O'
			}
		}
	}
}
