package problem0994

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type position struct{ row, col int }
	queue := []position{}
	amount := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == 2 {
				queue = append(queue, position{i, j})
			}
			if grid[i][j] == 1 {
				amount++
			}
		}
	}
	time := 0
	for len(queue) > 0 {
		nextQueue := []position{}
		for _, cur := range queue {
			row, col := cur.row, cur.col
			if row-1 > -1 && grid[row-1][col] == 1 {
				grid[row-1][col] = 2
				amount--
				nextQueue = append(nextQueue, position{row - 1, col})
			}
			if col+1 < n && grid[row][col+1] == 1 {
				grid[row][col+1] = 2
				amount--
				nextQueue = append(nextQueue, position{row, col + 1})
			}
			if row+1 < m && grid[row+1][col] == 1 {
				grid[row+1][col] = 2
				amount--
				nextQueue = append(nextQueue, position{row + 1, col})
			}
			if col-1 > -1 && grid[row][col-1] == 1 {
				grid[row][col-1] = 2
				amount--
				nextQueue = append(nextQueue, position{row, col - 1})
			}
		}
		if len(nextQueue) > 0 {
			time++
		}
		queue = nextQueue
	}
	if amount > 0 {
		return -1
	}
	return time
}
