package problem0045

// 45. 跳跃游戏 II

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。

// 贪心。
func jump(nums []int) int {
	n := len(nums)
	farthest, edge := 0, 0
	steps := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, nums[i]+i)
		if i == edge {
			edge = farthest
			steps++
		}
	}
	return steps
}

// 枚举。
func jump_enum(nums []int) int {
	n := len(nums)
	mins := make([]int, n)
	for i := 1; i < n; i++ {
		mins[i] = n
	}
	for i := 0; i < n-1; i++ {
		farthest := i + nums[i]
		for j := i + 1; j <= farthest && j < n; j++ {
			if mins[j] > mins[i]+1 {
				mins[j] = mins[i] + 1
			}
		}
	}
	return mins[n-1]
}
