package nc038

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}

	var top, right, bottom, left = 0, len(matrix[0]) - 1, len(matrix) - 1, 0
	amount := len(matrix) * len(matrix[0])
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
			amount--
		}

		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
			amount--
		}

		for i := right - 1; i >= left && amount > 0; i-- {
			result = append(result, matrix[bottom][i])
			amount--
		}

		for i := bottom - 1; i > top && amount > 0; i-- {
			result = append(result, matrix[i][left])
			amount--
		}

		top++
		right--
		bottom--
		left++
	}

	return result
}
