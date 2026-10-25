package problem0909

func snakesAndLadders(board [][]int) int {
	n := len(board)
	visited := make([]bool, n*n)
	var idToRowCol = func(id int) (int, int) {
		row, col := (id-1)/n, (id-1)%n
		if row%2 == 1 {
			col = n - 1 - col
		}
		row = n - 1 - row
		return row, col
	}
	type vertex struct{ id, step int }
	queue := []*vertex{{1, 0}}
	visited[0] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for i := 1; i <= 6; i++ {
			next := v.id + i
			row, col := idToRowCol(next)
			if board[row][col] != -1 {
				next = board[row][col]
			}
			if next == n*n {
				return v.step + 1
			}
			if !visited[next-1] {
				visited[next-1] = true
				queue = append(queue, &vertex{next, v.step + 1})
			}
		}
	}
	return -1
}
