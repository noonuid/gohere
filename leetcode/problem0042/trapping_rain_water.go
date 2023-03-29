package problem0042

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 先找到最大值的位置，再从两边向最大值靠拢计算雨水总量。
func trap(height []int) int {
	// 查找最大值的位置
	max := 0
	for i := 1; i < len(height); i++ {
		if height[i] > height[max] {
			max = i
		}
	}
	rain := 0
	// 最大值左边部分
	left, right := 0, 1
	for right < max {
		for height[right] < height[left] {
			rain += height[left] - height[right]
			right++
		}
		left = right
		right++
	}
	// 最大值右边部分
	left, right = len(height)-2, len(height)-1
	for left > max {
		for height[left] < height[right] {
			rain += height[right] - height[left]
			left--
		}
		right = left
		left--
	}
	return rain
}

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 使用双指针。
func trapTwoPointer(height []int) int {
	rain := 0
	left, right := 0, len(height)-1
	leftMaxHeight, rightMaxHeight := 0, 0
	for left < right {
		// 更新左右侧已遍历部分的最高高度
		if leftMaxHeight < height[left] {
			leftMaxHeight = height[left]
		}
		if rightMaxHeight < height[right] {
			rightMaxHeight = height[right]
		}
		if leftMaxHeight > rightMaxHeight {
			rain += rightMaxHeight - height[right]
			right--
		} else {
			rain += leftMaxHeight - height[left]
			left++
		}
	}
	return rain
}
