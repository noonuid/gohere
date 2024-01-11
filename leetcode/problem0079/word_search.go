package problem0079

// 79. 单词搜索

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 回溯法。
func exist(board [][]byte, word string) bool {
	m, n, length := len(board), len(board[0]), len(word)
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		if board[i][j] != word[index] {
			return false
		} else if index == length-1 {
			return true
		}
		used[i][j] = true
		if i-1 > -1 && !used[i-1][j] && dfs(i-1, j, index+1) {
			return true
		}
		if j+1 < n && !used[i][j+1] && dfs(i, j+1, index+1) {
			return true
		}
		if i+1 < m && !used[i+1][j] && dfs(i+1, j, index+1) {
			return true
		}
		if j-1 > -1 && !used[i][j-1] && dfs(i, j-1, index+1) {
			return true
		}
		used[i][j] = false
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
