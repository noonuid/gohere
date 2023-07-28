package problem0739

import "math"

// 给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
// 下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

// 单调栈。
func dailyTemperatures_stack(temperatures []int) []int {
	answers := make([]int, len(temperatures))
	stack := []int{0}
	for i := 1; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			preIndex := stack[len(stack)-1]
			answers[preIndex] = i - preIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answers
}

// 使用 next 数组存储当前温度之后的温度首次出现的索引。
func dailyTemperatures_next(temperatures []int) []int {
	answers := make([]int, len(temperatures))
	next := make([]int, 101)
	for i := 0; i < len(next); i++ {
		next[i] = math.MaxInt
	}
	// 反向遍历温度列表。
	for i := len(temperatures) - 1; i > -1; i-- {
		warmerIndex := math.MaxInt
		for temperature := temperatures[i] + 1; temperature < len(next); temperature++ {
			if next[temperature] < warmerIndex {
				warmerIndex = next[temperature]
			}
		}
		if warmerIndex < len(temperatures) {
			answers[i] = warmerIndex - i
		}
		next[temperatures[i]] = i
	}
	return answers
}

// 暴力枚举。
// 超出时间限制。
func dailyTemperatures_brute(temperatures []int) []int {
	answers := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		j := i + 1
		for j < len(temperatures) && temperatures[j] <= temperatures[i] {
			j++
		}
		if j < len(temperatures) {
			answers[i] = j - i
		}
	}
	return answers
}
