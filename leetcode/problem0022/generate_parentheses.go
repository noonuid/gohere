package problem0022

// 22. 括号生成

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 深度优先搜索加剪枝。
func generateParenthesis_dfs(n int) []string {
	res := []string{}
	var dfs func(path string, left, right int)
	dfs = func(path string, left, right int) {
		if left < 0 || left > right {
			return
		}
		if left == 0 && right == 0 {
			res = append(res, path)
		}
		dfs(path+"(", left-1, right)
		dfs(path+")", left, right-1)
	}
	dfs("", n, n)
	return res
}

// 暴力枚举。
func generateParenthesis_brute(n int) []string {
	res := []string{}
	isValid := func(path []byte) bool {
		stack := []byte{}
		for _, v := range path {
			if v == '(' {
				stack = append(stack, v)
			} else {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
		return len(stack) <= 0
	}
	var recursion func(path []byte)
	recursion = func(path []byte) {
		if len(path) == n*2 {
			if isValid(path) {
				res = append(res, string(path))
			}
			return
		}
		path = append(path, '(')
		recursion(path)
		path[len(path)-1] = ')'
		recursion(path)
	}
	recursion([]byte{})
	return res
}
