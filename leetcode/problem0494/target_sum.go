package problem0494

// 给你一个整数数组 nums 和一个整数 target 。

// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。

// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

// 回溯法。
func findTargetSumWays_backtrack(nums []int, target int) int {
	res := 0
	var backtrack func(cur int, sum int)
	backtrack = func(cur int, sum int) {
		if cur > len(nums)-1 {
			if sum == target {
				res++
			}
			return
		}
		backtrack(cur+1, sum+nums[cur])
		backtrack(cur+1, sum-nums[cur])
	}
	backtrack(0, 0)
	return res
}

// 动态规划。
func findTargetSumWays_dynamic_programming(nums []int, target int) int {
	// 求 nums 中所有元素的和。
	sum := 0
	for _, num := range nums {
		sum += num
	}
	diff := sum - target
	// diff 为负数或者奇数时，找不到符合条件的表达式。
	if diff < 0 || diff%2 == 1 {
		return 0
	}
	neg := diff / 2
	dp := make([]int, neg+1)
	dp[0] = 1
	for _, num := range nums {
		// 使用滚动数组的方式，内层循环需采用倒序遍历的方式，利用上一轮计算的值获取本轮的值。
		for j := neg; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[neg]
}
