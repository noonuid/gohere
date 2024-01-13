package problem0200

// 200. 岛屿数量

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

// 此外，你可以假设该网格的四条边均被水包围。

// 广度优先搜索。
// 超出内存限制。
func numIslands_bfs(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	type info struct{ row, col int }
	num := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				num++
				queue := []info{{row, col}}
				for len(queue) > 0 {
					i, j := queue[0].row, queue[0].col
					grid[i][j] = '0'
					queue = queue[1:]
					if i-1 > -1 && grid[i-1][j] == '1' {
						queue = append(queue, info{i - 1, j})
					}
					if j+1 < cols && grid[i][j+1] == '1' {
						queue = append(queue, info{i, j + 1})
					}
					if i+1 < rows && grid[i+1][j] == '1' {
						queue = append(queue, info{i + 1, j})
					}
					if j-1 > -1 && grid[i][j-1] == '1' {
						queue = append(queue, info{i, j - 1})
					}
				}
			}
		}
	}
	return num
}

// 深度优先搜索。
func numIslands_dfs(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		if row-1 > -1 {
			dfs(row-1, col)
		}
		if col+1 < cols {
			dfs(row, col+1)
		}
		if row+1 < rows {
			dfs(row+1, col)
		}
		if col-1 > -1 {
			dfs(row, col-1)
		}
	}
	num := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				num++
				dfs(row, col)
			}
		}
	}
	return num
}
