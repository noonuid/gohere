package problem0209

// 滑动窗口解法
func minSubArrayLen(target int, nums []int) int {
	// 分别代表子数组的前后两个位置
	var front = 0
	var back = 0

	// 记录当前子数组的和
	var sum = 0
	var minLength = len(nums) + 1
	for front < len(nums) {
		sum += nums[front]

		// 从子数组的后面丢弃元素
		for back <= front && sum >= target {
			if minLength > front-back+1 {
				minLength = front - back + 1
			}

			sum -= nums[back]
			back++
		}

		front++
	}

	if minLength <= len(nums) {
		return minLength
	}
	return 0
}
