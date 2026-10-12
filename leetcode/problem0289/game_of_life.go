package problem0289

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	around := make([][]int, m)
	for i := 0; i < m; i++ {
		around[i] = make([]int, n)
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if row-1 > -1 && col-1 > -1 && board[row-1][col-1] == 1 {
				around[row][col]++
			}
			if row-1 > -1 && board[row-1][col] == 1 {
				around[row][col]++
			}
			if row-1 > -1 && col+1 < n && board[row-1][col+1] == 1 {
				around[row][col]++
			}
			if col+1 < n && board[row][col+1] == 1 {
				around[row][col]++
			}
			if row+1 < m && col+1 < n && board[row+1][col+1] == 1 {
				around[row][col]++
			}
			if row+1 < m && board[row+1][col] == 1 {
				around[row][col]++
			}
			if row+1 < m && col-1 > -1 && board[row+1][col-1] == 1 {
				around[row][col]++
			}
			if col-1 > -1 && board[row][col-1] == 1 {
				around[row][col]++
			}
		}
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == 1 && (around[row][col] < 2 || around[row][col] > 3) {
				board[row][col] = 0
			}
			if board[row][col] == 0 && around[row][col] == 3 {
				board[row][col] = 1
			}
		}
	}
}

func gameOfLife_composite_state(board [][]int) {
	m, n := len(board), len(board[0])
	neighbors := []int{-1, 0, 1}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			live := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !(neighbors[i] == 0 && neighbors[j] == 0) {
						r, c := row+neighbors[i], col+neighbors[j]
						if (-1 < r && r < m) && (-1 < c && c < n) && (board[r][c] == 1 || board[r][c] == -1) {
							live++
						}
					}
				}
			}
			if board[row][col] == 1 && (live < 2 || live > 3) {
				board[row][col] = -1
			}
			if board[row][col] == 0 && live == 3 {
				board[row][col] = 2
			}
		}
	}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] > 0 {
				board[row][col] = 1
			} else {
				board[row][col] = 0
			}
		}
	}
}
