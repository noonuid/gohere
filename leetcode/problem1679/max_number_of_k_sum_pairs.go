package problem1679

import "slices"

func maxOperations(nums []int, k int) int {
	count := 0
	slices.Sort(nums)
	left, right := 0, len(nums)-1
	for left < right {
		switch sum := nums[left] + nums[right]; {
		case sum < k:
			left++
		case sum > k:
			right--
		case sum == k:
			count++
			left, right = left+1, right-1
		}
	}
	return count
}

func maxOperations_map(nums []int, k int) int {
	count := 0
	m := make(map[int]int)
	for _, num := range nums {
		if m[k-num] > 0 {
			count++
			m[k-num]--
		} else {
			m[num]++
		}
	}
	return count
}
