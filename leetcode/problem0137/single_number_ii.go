package problem0137

import "sort"

// 137. 只出现一次的数字 II

// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

// 你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。

// 位运算。
func singleNumber_bitwise(nums []int) int {
	answer := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 == 1 {
			answer |= 1 << i
		}
	}
	return int(answer)
}

// 排序。
func singleNumber_sort(nums []int) int {
	sort.Ints(nums)
	for i := 0; i <= len(nums)-4; i = i + 3 {
		if nums[i] != nums[i+2] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
