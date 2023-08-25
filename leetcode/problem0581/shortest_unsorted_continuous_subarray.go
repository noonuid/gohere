package problem0581

import "sort"

// 581. 最短无序连续子数组

// 给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

// 请你找出符合题意的 最短 子数组，并输出它的长度。

// 在遍历过程中确定连续子数组的左右边界。
func findUnsortedSubarray_traversal(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	// left，righ 分别表示连续子数组的左右边界。
	left, right := n-1, 0
	min, max := nums[n-1], nums[0]
	for i := 1; i < n; i++ {
		if nums[i] < max {
			right = i
		} else {
			max = nums[i]
		}

		if nums[n-1-i] > min {
			left = n - 1 - i
		} else {
			min = nums[n-1-i]
		}
	}
	if right > left {
		return right - left + 1
	} else {
		return 0
	}
}

// 先将原数组排序，再从两头向中间逐个数字比对，找到子数组的两端。
func findUnsortedSubarray_sort(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	numsCopy := make([]int, n)
	copy(numsCopy, nums)
	sort.Ints(numsCopy)
	begin, end := 0, n-1
	for begin < n && nums[begin] == numsCopy[begin] {
		begin++
	}
	if begin < end {
		for end > -1 && nums[end] == numsCopy[end] {
			end--
		}
		return end - begin + 1
	}
	return 0
}
