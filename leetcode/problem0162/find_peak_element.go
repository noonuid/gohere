package problem0162

import "math"

// 162. 寻找峰值

// 峰值元素是指其值严格大于左右相邻值的元素。

// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

// 你可以假设 nums[-1] = nums[n] = -∞ 。

// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

// 遍历找寻最大值。
func findPeakElement_max(nums []int) int {
	index := 0
	for i, num := range nums {
		if num > nums[index] {
			index = i
		}
	}
	return index
}

// 二分查找。
func findPeakElement_binary(nums []int) int {
	n := len(nums)

	get := func(index int) int {
		if index == -1 || index == n {
			return math.MinInt
		}
		return nums[index]
	}

	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		} else if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
