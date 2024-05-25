package problem0200

import "container/list"

// 200. 岛屿数量

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

// 此外，你可以假设该网格的四条边均被水包围。

// 广度优先搜索。
func numIslands_bfs(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	num := 0
	type pos struct{ row, col int }
	queue := list.New()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				num++

				queue.PushBack(pos{i, j})
				grid[i][j] = '0'
				for queue.Len() > 0 {
					front := queue.Front()
					row, col := front.Value.(pos).row, front.Value.(pos).col
					queue.Remove(front)
					if row-1 > -1 && grid[row-1][col] == '1' {
						grid[row-1][col] = '0'
						queue.PushBack(pos{row - 1, col})
					}
					if col+1 < n && grid[row][col+1] == '1' {
						grid[row][col+1] = '0'
						queue.PushBack(pos{row, col + 1})
					}
					if row+1 < m && grid[row+1][col] == '1' {
						grid[row+1][col] = '0'
						queue.PushBack(pos{row + 1, col})
					}
					if col-1 > -1 && grid[row][col-1] == '1' {
						grid[row][col-1] = '0'
						queue.PushBack(pos{row, col - 1})
					}
				}
			}
		}
	}
	return num
}

// 深度优先搜索。
func numIslands_dfs(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= m || col >= n || grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row+1, col)
		dfs(row, col-1)
	}
	num := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				num++
				dfs(i, j)
			}
		}
	}
	return num
}
