package problem0212

func findWords(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])

	type TrieNode struct {
		children [26]*TrieNode
		word     string
	}

	root := &TrieNode{}

	insert := func(word string) {
		node := root
		for i := 0; i < len(word); i++ {
			childIndex := word[i] - 'a'
			if node.children[childIndex] == nil {
				node.children[childIndex] = &TrieNode{}
			}
			node = node.children[childIndex]
		}
		node.word = word
	}

	for _, word := range words {
		insert(word)
	}

	seen := make(map[string]bool)

	var dfs func(row, col int, node *TrieNode)
	dfs = func(row, col int, node *TrieNode) {
		if !(-1 < row && row < m && -1 < col && col < n) || board[row][col] == '#' {
			return
		}

		ch := board[row][col]
		child := node.children[ch-'a']
		if child == nil {
			return
		}

		if child.word != "" {
			seen[child.word] = true
		}

		board[row][col] = '#'
		dfs(row-1, col, child)
		dfs(row+1, col, child)
		dfs(row, col-1, child)
		dfs(row, col+1, child)
		board[row][col] = ch
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}

	result := make([]string, 0, len(seen))
	for word := range seen {
		result = append(result, word)
	}
	return result
}

func findWords_delete_matched_words(board [][]byte, words []string) []string {
	m, n := len(board), len(board[0])

	type TrieNode struct {
		children map[byte]*TrieNode
		word     string
	}

	root := &TrieNode{children: map[byte]*TrieNode{}}

	insert := func(word string) {
		node := root
		for i := 0; i < len(word); i++ {
			if node.children[word[i]] == nil {
				node.children[word[i]] = &TrieNode{children: map[byte]*TrieNode{}}
			}
			node = node.children[word[i]]
		}
		node.word = word
	}

	for _, word := range words {
		insert(word)
	}

	result := []string{}

	var dfs func(row, col int, node *TrieNode)
	dfs = func(row, col int, node *TrieNode) {
		if !(-1 < row && row < m && -1 < col && col < n) || board[row][col] == '#' {
			return
		}

		ch := board[row][col]
		child := node.children[ch]
		if child == nil {
			return
		}

		if child.word != "" {
			result = append(result, child.word)
			child.word = ""
		}

		if len(child.children) > 0 {
			board[row][col] = '#'
			dfs(row-1, col, child)
			dfs(row+1, col, child)
			dfs(row, col-1, child)
			dfs(row, col+1, child)
			board[row][col] = ch
		}

		if len(child.children) == 0 {
			delete(node.children, ch)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}

	return result
}
