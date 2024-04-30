package problem0324

import "sort"

// 324. 摆动排序 II

// 给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

// 你可以假设所有输入数组都可以得到满足题目要求的结果。

// 先升序排序，再将前后段元素按倒序依次交叉合并。
func wiggleSort(nums []int) {
	n := len(nums)
	numsCopy := make([]int, n)
	copy(numsCopy, nums)
	sort.Ints(numsCopy)
	mid := (n - 1) / 2
	for i, j, cur := mid, n-1, 0; j > mid; i, j, cur = i-1, j-1, cur+2 {
		nums[cur], nums[cur+1] = numsCopy[i], numsCopy[j]
	}
	if n%2 == 1 {
		nums[n-1] = numsCopy[0]
	}
}
