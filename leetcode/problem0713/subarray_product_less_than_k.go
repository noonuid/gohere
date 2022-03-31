package problem0713

func numSubarrayProductLessThanK(nums []int, k int) int {
	var result = 0

	var left, right = 0, 0
	var product = 1
	for right < len(nums) {
		product *= nums[right]

		for left <= right && product >= k {
			product /= nums[left]
			left++
		}

		if left <= right {
			result += right - left + 1
		}

		right++
	}

	return result
}
