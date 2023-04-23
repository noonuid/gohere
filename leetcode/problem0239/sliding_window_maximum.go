package problem0239

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

// 返回 滑动窗口中的最大值 。

// 单调队列。
func maxSlidingWindow(nums []int, k int) []int {
	maxes := make([]int, len(nums)-k+1)
	// queue 中存储索引，索引对应的值严格单调递减。
	queue := make([]int, 0, k)
	// 将窗口右边的元素的下标加入到 queue 中。
	push := func(right int) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[right] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, right)
	}
	// 建立窗口。
	for right := 0; right < k; right++ {
		push(right)
	}
	maxes[0] = nums[queue[0]]
	// 移动窗口。
	for left, right := 1, k; right < len(nums); left, right = left+1, right+1 {
		// 如果队首元素此时在窗口之外，则从队列中移除其下标。
		if queue[0] < left {
			queue = queue[1:]
		}
		push(right)
		maxes[left] = nums[queue[0]]
	}
	return maxes
}

// 暴力枚举。超出时间限制。
func maxSlidingWindow_brute_force(nums []int, k int) []int {
	getMaxIndex := func(window []int) int {
		max := 0
		for i := 1; i < len(window); i++ {
			if window[max] < window[i] {
				max = i
			}
		}
		return max
	}
	maxes := make([]int, len(nums)-k+1)
	maxIndex := getMaxIndex(nums[:k])
	maxes[0] = nums[maxIndex]
	for left, right := 1, k; left < len(nums)-k+1; left++ {
		right = left + k - 1
		if nums[right] >= nums[maxIndex] {
			maxIndex = right
		} else if maxIndex < left {
			maxIndex = left + getMaxIndex(nums[left:left+k])
		}
		maxes[left] = nums[maxIndex]
	}
	return maxes
}
