package problem0077

// 77. 组合

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

// 你可以按 任何顺序 返回答案。

func combine(n int, k int) [][]int {
	numerator, denominator := 1, 1
	for i := 0; i < k; i++ {
		numerator *= n - i
		denominator *= (i + 1)
	}
	result := make([][]int, 0, numerator/denominator)
	path := make([]int, 0, k)

	var backtrack func(cur int)
	backtrack = func(cur int) {
		if len(path) == k {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		need := k - len(path)
		for i := cur; need <= n-i+1; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)
	return result
}
