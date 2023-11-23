package problem0034

// 34. 在排序数组中查找元素的第一个和最后一个位置

// 给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

// 如果数组中不存在目标值 target，返回 [-1, -1]。

// 你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

// 二分查找。
func searchRange(nums []int, target int) []int {
	n := len(nums)
	binarySearch := func(leftMost bool) int {
		left, right, pos := 0, n-1, -1
		for left <= right {
			mid := (left + right) / 2
			if cur := nums[mid]; cur < target || (!leftMost && cur <= target) {
				left = mid + 1
				pos = mid
			} else {
				right = mid - 1
			}
		}
		return pos
	}
	// 数组中最后一个小于 target 的元素的位置。
	lastLess := binarySearch(true)
	if firstEqual := lastLess + 1; -1 < firstEqual && firstEqual < n && nums[firstEqual] == target {
		// 数组中最后一个等于 target 的元素的位置。
		lastEqual := binarySearch(false)
		return []int{firstEqual, lastEqual}
	} else {
		return []int{-1, -1}
	}
}
