package offer03

// 剑指 Offer 03. 数组中重复的数字

// 找出数组中重复的数字。

// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
// 数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
// 请找出数组中任意一个重复的数字。

// 哈希表。
func findRepeatNumber_hash(nums []int) int {
	l := len(nums)
	m := make(map[int]struct{})
	for i := 0; i < l; i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = struct{}{}
		} else {
			return nums[i]
		}
	}
	return 0
}

// 暴力枚举。
func findRepeatNumber_brute(nums []int) int {
	l := len(nums)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}
