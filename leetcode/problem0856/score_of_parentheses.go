package problem0856

// 方法一：计算得分时只考虑“()”项，而将“()”项外层的括号作为增加内层“()”项深度的中间项，不直接计算其得分
func scoreOfParentheses(s string) int {
	// 初始分数权重为 1，整体得分为 0
	var weight, score = 1, 0

	for index, ch := range s {
		if ch == '(' {
			// 当前位置的“()”项的深度增加，分数权重增加
			weight *= 2
		} else {
			// 当前位置的“()”项的深度减少，分数权重减少
			weight /= 2

			// 遇到“()”项，计算一次得分
			if s[index-1] == '(' {
				score += weight
			}
		}
	}

	return score
}

// 方法二：使用栈存储 '(' 和中间分数
func scoreOfParenthesesStack(s string) int {
	var stack = make([]int, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 遇到 '('，相当于分数 0，将 0 入栈
			stack = append(stack, 0)
		} else {
			// 遇到 ')'，两种情况：“()”和“(ABC)”
			if stack[len(stack)-1] == 0 {
				// 当栈顶为 '('，为“()”的情况，分数为 1
				stack = stack[:len(stack)-1]
				stack = append(stack, 1)
			} else {
				// 当栈顶为子项得分，为“(ABC)”的情况，分数为 (A+B+C)*2
				var scopeScore = 0
				for stack[len(stack)-1] > 0 {
					scopeScore += stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				// 将对应的 '(' 出栈
				stack = stack[:len(stack)-1]
				stack = append(stack, scopeScore*2)
			}
		}
	}

	// 栈中剩下的为最外层各个子项的得分，属于 ABC 的情况
	var score = 0
	for len(stack) > 0 {
		score += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return score
}
