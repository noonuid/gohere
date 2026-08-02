package problem0334

import "math"

func increasingTriplet(nums []int) bool {
	n := len(nums)
	leftMin, rightMax := make([]int, n), make([]int, n)
	leftMin[0], rightMax[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
		rightMax[n-1-i] = max(nums[n-1-i], rightMax[n-i])
	}
	for i := 1; i < n-1; i++ {
		if leftMin[i-1] < nums[i] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}

func increasingTriplet_greedy(nums []int) bool {
	n := len(nums)
	first, second := nums[0], math.MaxInt
	for i := 1; i < n; i++ {
		if second < nums[i] {
			return true
		} else if first < nums[i] {
			second = nums[i]
		} else {
			first = nums[i]
		}
	}
	return false
}
