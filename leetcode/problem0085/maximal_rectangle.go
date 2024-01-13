package problem0085

// 85. 最大矩形

// 给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

// 先使用动态规划得到一系列柱状图，再使用单调栈计算每个柱状图中最大矩形的面积，
// 这些面积中的最大值就是最终的答案。
func maximalRectangle(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	maxArea := 0
	histograms := make([][]int, rows)
	for row := 0; row < rows; row++ {
		// 第一个和最后一个元素 0 作为每个柱状图头部和尾部的哨兵。
		histograms[row] = make([]int, cols+2)
		for col := 1; col < cols+1; col++ {
			if matrix[row][col-1] == '1' {
				if row == 0 {
					histograms[0][col] = 1
				} else {
					histograms[row][col] = histograms[row-1][col] + 1
				}
			}
		}
	}
	// 对每一个柱状图计算其中最大的矩形面积。
	for row := 0; row < rows; row++ {
		heights := histograms[row]
		stack := make([]int, 1, cols+1)
		// 这个 0 代表柱状图头部的哨兵。
		// stack[0] = 0
		for col := 1; col < cols+2; col++ {
			for heights[stack[len(stack)-1]] > heights[col] {
				height := heights[stack[len(stack)-1]]
				for heights[stack[len(stack)-1]] == height {
					stack = stack[:len(stack)-1]
				}
				width := col - stack[len(stack)-1] - 1
				if area := height * width; maxArea < area {
					maxArea = area
				}
			}
			stack = append(stack, col)
		}
	}
	return maxArea
}
