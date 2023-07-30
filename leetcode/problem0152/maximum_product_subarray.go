package problem0152

import "math"

// 152. 乘积最大子数组

// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

// 测试用例的答案是一个 32-位 整数。

// 子数组 是数组的连续子序列。

func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

// 动态规划。
func maxProduct_dynamic_programming(nums []int) int {
	// 滚动数组思想，只使用两个变量维护状态。
	res, dp_max, dp_min := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num, m, n := nums[i], dp_max, dp_min
		dp_max = max(m*num, n*num, num)
		dp_min = min(m*num, n*num, num)
		if dp_max > res {
			res = dp_max
		}
	}
	return res
}

// 暴力枚举。
// 超出时间限制。
func maxProduct_brute(nums []int) int {
	max := math.MinInt
	for left := 0; left < len(nums); left++ {
		for right := left; right < len(nums); right++ {
			product := 1
			for i := left; i <= right; i++ {
				product *= nums[i]
			}
			if product > max {
				max = product
			}
		}
	}
	return max
}
