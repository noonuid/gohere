package problem0300

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

// 动态规划。
func lengthOfLIS(nums []int) int {
	// f[i] 表示以 f[i] 结尾的最长严格递增子序列的长度。
	f := make([]int, len(nums))
	max := 1
	for i := 0; i < len(f); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && f[i] < f[j]+1 {
				f[i] = f[j] + 1
			}
		}
		if max < f[i] {
			max = f[i]
		}
	}
	return max
}
