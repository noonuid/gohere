package problem0036

// 36. 有效的数独

// 请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）

// 注意：

// 一个有效的数独（部分已被填充）不一定是可解的。
// 只需要根据以上规则，验证已经填入的数字是否有效即可。
// 空白格用 '.' 表示。

// 一次遍历，使用数组存储当前遍历的数字在同一行、同一列或者同一个九宫格中是否已经出现过。
func isValidSudoku(board [][]byte) bool {
	rows := new([9][9]bool)
	cols := new([9][9]bool)
	smalls := new([3][3][9]bool)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				index := board[row][col] - '0' - 1
				if rows[row][index] || cols[col][index] || smalls[row/3][col/3][index] {
					return false
				}
				rows[row][index] = true
				cols[col][index] = true
				smalls[row/3][col/3][index] = true
			}
		}
	}
	return true
}

// 穷举法。
func isValidSudoku_brute(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		m := make(map[byte]struct{}, 9)
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				if _, ok := m[board[row][col]]; ok {
					return false
				} else {
					m[board[row][col]] = struct{}{}
				}
			}
		}
	}
	for col := 0; col < 9; col++ {
		m := make(map[byte]struct{}, 9)
		for row := 0; row < 9; row++ {
			if board[row][col] != '.' {
				if _, ok := m[board[row][col]]; ok {
					return false
				} else {
					m[board[row][col]] = struct{}{}
				}
			}
		}
	}
	for small := 0; small < 9; small++ {
		m := make(map[byte]struct{}, 9)
		for id := 0; id < 9; id++ {
			row, col := (small/3)*3+id/3, (small%3)*3+id%3
			if board[row][col] != '.' {
				if _, ok := m[board[row][col]]; ok {
					return false
				} else {
					m[board[row][col]] = struct{}{}
				}
			}
		}
	}
	return true
}
