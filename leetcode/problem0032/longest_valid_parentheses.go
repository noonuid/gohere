package problem0032

// 32. 最长有效括号

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

// 两次遍历。
func longestValidParentheses_double_traversals(s string) int {
	length := len(s)
	if length < 2 {
		return 0
	}
	longest := 0
	left, right := 0, 0
	for i := 0; i < length; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left < right {
			left, right = 0, 0
		} else if left == right && longest < left+right {
			longest = left + right
		}
	}
	left, right = 0, 0
	for i := length - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left > right {
			left, right = 0, 0
		} else if left == right && longest < left+right {
			longest = left + right
		}
	}
	return longest
}

func longestValidParentheses_stack(s string) int {
	length := len(s)
	if length < 2 {
		return 0
	}
	longest := 0
	stack := make([]int, 1, length+1)
	stack[0] = -1
	for i := 0; i < length; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				// 当前的 ')' 为已遍历元素中最后一个未匹配的右括号。
				stack = append(stack, i)
			} else if subLen := i - stack[len(stack)-1]; longest < subLen {
				longest = subLen
			}
		}
	}
	return longest
}

// 动态规划。
func longestValidParentheses_dynamic_programming(s string) int {
	length := len(s)
	if length < 2 {
		return 0
	}
	longest := 0
	f := make([]int, length)
	if s[0] == '(' && s[1] == ')' {
		f[1], longest = 2, 2
	}
	for i := 2; i < length; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { // s[:i+1] 为 "...()" 的形式。
				f[i] = f[i-2] + 2
			} else if (i-1)-f[i-1] > -1 && s[(i-1)-f[i-1]] == '(' {
				// s[:i+1] 为 "...))" 的形式，并且 i 处的 ')' 能够匹配到一个 '('。

				if subLen := f[i-1] + 2; i-subLen > -1 {
					f[i] = subLen + f[i-subLen]
				} else {
					f[i] = subLen
				}
			}

			if longest < f[i] {
				longest = f[i]
			}
		}
	}
	return longest
}

// 暴力枚举。
// 超出时间限制。
func longestValidParentheses_brute(s string) int {
	length := len(s)
	longest := 0
	valid := func(str string) bool {
		balance := 0
		for _, r := range str {
			if r == '(' {
				balance++
			} else {
				if balance < 1 {
					return false
				}
				balance--
			}
		}
		return balance == 0
	}
	for left := 0; left < length-1; left++ {
		right := left + 2
		if longest > 2 {
			right = left + longest
		}
		for ; right <= length; right = right + 2 {
			if valid(s[left:right]) {
				longest = right - left
			}
		}
	}
	return longest
}
