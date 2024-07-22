package problem0918

import "math"

// 将「最大子数组和」问题转化为「最小子数组和」问题。
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	preMax, preMin := nums[0], nums[0]
	maxSubSum, minSubSum := nums[0], nums[0]
	sum := nums[0]
	for i := 1; i < n; i++ {
		preMax = max(preMax+nums[i], nums[i])
		maxSubSum = max(maxSubSum, preMax)
		preMin = min(preMin+nums[i], nums[i])
		minSubSum = min(minSubSum, preMin)
		sum += nums[i]
	}
	if maxSubSum < 0 {
		return maxSubSum
	} else {
		return max(maxSubSum, sum-minSubSum)
	}
}

// 动态规划。
func maxSubarraySumCircular_dynamic_programming(nums []int) int {
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	preMax := nums[0]
	leftMax, leftSum := make([]int, n), nums[0]
	leftMax[0] = nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		preMax = max(preMax+nums[i], nums[i])
		res = max(res, preMax)
		leftSum += nums[i]
		leftMax[i] = max(leftMax[i-1], leftSum)
	}
	rightSum := 0
	for i := n - 1; i > 0; i-- {
		rightSum += nums[i]
		res = max(res, leftMax[i-1]+rightSum)
	}
	return res
}

// 枚举。
// 超出时间限制。
func maxSubarraySumCircular_enum(nums []int) int {
	n := len(nums)
	sum := func(start, end int) int {
		result := 0
		if start <= end {
			for i := start; i <= end; i++ {
				result += nums[i]
			}
		} else {
			for i := start; i < n; i++ {
				result += nums[i]
			}
			for i := 0; i <= end; i++ {
				result += nums[i]
			}
		}
		return result
	}
	max := math.MinInt
	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			if s := sum(start, end); max < s {
				max = s
			}
		}
		for end := 0; end < start; end++ {
			if s := sum(start, end); max < s {
				max = s
			}
		}
	}
	return max
}
