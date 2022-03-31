package problem0053

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	addSum := 0
	for i := 0; i < len(nums); i++ {
		addSum += nums[i]
		if addSum > maxSum {
			maxSum = addSum
		}
		if addSum < 0 {
			addSum = 0
		}
	}
	return maxSum
}
