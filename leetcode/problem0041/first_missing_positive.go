package problem0041

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range n {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if num < n+1 && nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return n + 1
}

func firstMissingPositive_swap(nums []int) int {
	n := len(nums)
	for i := range n {
		for 0 < nums[i] && nums[i] < n+1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range n {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
