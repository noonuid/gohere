package problem0052

func totalNQueens(n int) int {
	result := 0
	cols := make([]bool, n)
	diags1 := make([]bool, 2*n-1)
	diags2 := make([]bool, 2*n-1)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			result++
			return
		}

		for col := 0; col < n; col++ {
			index1, index2 := row-col+n-1, row+col
			if !cols[col] && !diags1[index1] && !diags2[index2] {
				cols[col], diags1[index1], diags2[index2] = true, true, true
				backtrack(row + 1)
				cols[col], diags1[index1], diags2[index2] = false, false, false
			}
		}
	}

	backtrack(0)
	return result
}

func totalNQueens_bitwise_operation(n int) int {
	result := 0

	var backtrack func(row, cols, diags1, diags2 int)
	backtrack = func(row, cols, diags1, diags2 int) {
		if row == n {
			result++
			return
		}

		availablePositions := (1<<n - 1) &^ (cols | diags1 | diags2)
		for availablePositions > 0 {
			position := availablePositions & (-availablePositions)
			backtrack(row+1, cols|position, (diags1|position)<<1, (diags2|position)>>1)
			availablePositions = availablePositions &^ position
		}
	}

	backtrack(0, 0, 0, 0)
	return result
}
