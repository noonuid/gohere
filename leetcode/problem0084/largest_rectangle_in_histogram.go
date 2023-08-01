package problem0084

// 84. 柱状图中最大的矩形

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

// 单调栈。
func largestRectangleArea_monotone_stack(heights []int) int {
	res, length := 0, len(heights)
	// stack 为单调栈，存储 heights 的索引。
	stack := make([]int, 0, length)
	for i := 0; i < length; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			// 进入循环表示以栈顶元素对应高度为高度的矩形的最大宽度可以确定了。
			height, width := heights[stack[len(stack)-1]], 0
			stack = stack[:len(stack)-1]
			// 若栈内元素对应的高度是相同的，则以这些高度为高度的矩形的最大宽度是一样的，可以直接出栈。
			for len(stack) > 0 && heights[stack[len(stack)-1]] == height {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			if area := height * width; area > res {
				res = area
			}
		}
		stack = append(stack, i)
	}
	// heights 遍历完成后，栈内剩余元素对应矩形的最大宽度都可以确定了。
	for len(stack) > 0 {
		height, width := heights[stack[len(stack)-1]], 0
		stack = stack[:len(stack)-1]
		for len(stack) > 0 && heights[stack[len(stack)-1]] == height {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			width = length
		} else {
			width = length - stack[len(stack)-1] - 1
		}
		if area := height * width; area > res {
			res = area
		}
	}
	return res
}

// 单调栈加哨兵。
func largestRectangleArea_monotone_stack_sentinel(heights []int) int {
	// 改造后的高度数组。
	hs := make([]int, len(heights)+2)
	res, length := 0, len(hs)
	copy(hs[1:length-1], heights)
	// 头尾设置哨兵。
	// hs[0], hs[length-1] = 0, 0
	stack := make([]int, 1, length)
	stack[0] = 0
	for i := 1; i < length; i++ {
		for len(stack) > 0 && hs[stack[len(stack)-1]] > hs[i] {
			// 进入循环表示以栈顶元素对应高度为高度的矩形的最大宽度可以确定了。
			height := hs[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			// 若栈内元素对应的高度是相同的，则以这些高度为高度的矩形的最大宽度是一样的，可以直接出栈。
			for len(stack) > 0 && hs[stack[len(stack)-1]] == height {
				stack = stack[:len(stack)-1]
			}
			width := i - stack[len(stack)-1] - 1
			if area := height * width; area > res {
				res = area
			}
		}
		stack = append(stack, i)
	}
	return res
}

// 以每个柱形的高度作为矩形的高度进行暴力枚举。
// 超出时间限制。
func largestRectangleArea_brute_height(heights []int) int {
	res := 0
	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for left-1 > -1 && heights[left-1] >= heights[i] {
			left--
		}
		for right+1 < len(heights) && heights[right+1] >= heights[i] {
			right++
		}
		if cur := (right - left + 1) * heights[i]; cur > res {
			res = cur
		}
	}
	return res
}

// 暴力枚举矩形左右两边的位置。
// 超出时间限制。
func largestRectangleArea_brute_left_right(heights []int) int {
	res := 0
	for left := 0; left < len(heights); left++ {
		for right := left; right < len(heights); right++ {
			height := heights[left]
			for i := left + 1; i <= right; i++ {
				if height > heights[i] {
					height = heights[i]
				}
			}
			if cur := height * (right - left + 1); cur > res {
				res = cur
			}
		}
	}
	return res
}
