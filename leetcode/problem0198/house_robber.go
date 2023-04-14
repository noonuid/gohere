package problem0198

// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
// 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

// 动态规划。
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// f[i] 代表 i 及 i 之前的所有房屋能够获得的最大金额。
	f := make([]int, len(nums))
	f[0], f[1] = nums[0], nums[0]
	if f[1] < nums[1] {
		f[1] = nums[1]
	}
	for i := 2; i < len(f); i++ {
		f[i] = f[i-2] + nums[i]
		if f[i] < f[i-1] {
			f[i] = f[i-1]
		}
	}
	return f[len(nums)-1]
}
