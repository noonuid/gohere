package problem0416

import "math"

// 416. 分割等和子集416. 分割等和子集

// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

// 动态规划。
func canPartition_dynamic_programming(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, maxNum := 0, math.MinInt
	for _, num := range nums {
		sum += num
		if maxNum < num {
			maxNum = num
		}
	}
	target := sum / 2
	if sum%2 == 1 || maxNum > target {
		return false
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		for j := 1; j <= target; j++ {
			if nums[i] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}

// 暴力枚举。
// 超出时间限制。
func canPartition_brute(nums []int) bool {
	n := len(nums)
	var can func(i, sum1, sum2 int) bool
	can = func(i, sum1, sum2 int) bool {
		if i == n {
			return sum1 == sum2
		}

		return can(i+1, sum1+nums[i], sum2) || can(i+1, sum1, sum2+nums[i])
	}
	return can(0, 0, 0)
}
