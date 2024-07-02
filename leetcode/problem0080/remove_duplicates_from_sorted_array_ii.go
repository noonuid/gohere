package problem0080

// 80. 删除有序数组中的重复项 II

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。

// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

// 双指针。
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	i := 1
	for j := 2; j < n; j++ {
		if (nums[i] == nums[j] && nums[i-1] < nums[i]) || nums[i] < nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
