package problem0301

// 301. 删除无效的括号

// 给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。

// 返回所有可能的结果。答案可以按 任意顺序 返回。

// 回溯法。
func removeInvalidParentheses_backtracking(s string) []string {
	answer := []string{}
	valid := func(str string) bool {
		left := 0
		for _, ch := range str {
			if ch == '(' {
				left++
			} else if ch == ')' {
				left--
				if left < 0 {
					return false
				}
			}
		}
		return left == 0
	}
	var dfs func(str string, start, leftRm, rightRm int)
	dfs = func(str string, start, leftRm, rightRm int) {
		if leftRm+rightRm == 0 {
			if valid(str) {
				answer = append(answer, str)
			}
			return
		}
		for i := start; i < len(str); i++ {
			if len(str)-i < leftRm+rightRm {
				return
			}
			if i > start && str[i] == str[i-1] {
				continue
			}
			if leftRm > 0 && str[i] == '(' {
				dfs(str[:i]+str[i+1:], i, leftRm-1, rightRm)
			} else if rightRm > 0 && str[i] == ')' {
				dfs(str[:i]+str[i+1:], i, leftRm, rightRm-1)
			}
		}
	}
	leftRm, rightRm := 0, 0
	for _, ch := range s {
		if ch == '(' {
			leftRm++
		} else if ch == ')' {
			if leftRm > 0 {
				leftRm--
			} else {
				rightRm++
			}
		}
	}
	dfs(s, 0, leftRm, rightRm)
	return answer
}

// 广度优先搜索。
func removeInvalidParentheses_bfs(s string) []string {
	answer := []string{}
	valid := func(str string) bool {
		left := 0
		for _, ch := range str {
			if ch == '(' {
				left++
			} else if ch == ')' {
				left--
				if left < 0 {
					return false
				}
			}
		}
		return left == 0
	}
	queue := map[string]struct{}{s: {}}
	for {
		for str := range queue {
			if valid(str) {
				answer = append(answer, str)
			}
		}
		if len(answer) > 0 {
			return answer
		}
		nextQueue := map[string]struct{}{}
		for str := range queue {
			for i := 0; i < len(str); i++ {
				if i > 0 && str[i] == str[i-1] {
					continue
				}
				if str[i] == '(' || str[i] == ')' {
					nextQueue[str[:i]+str[i+1:]] = struct{}{}
				}
			}
		}
		queue = nextQueue
	}
}
