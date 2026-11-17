package problem0120

import "math"

func minimumTotal(triangle [][]int) int {
	rowNum := len(triangle)
	dp := make([]int, len(triangle[rowNum-1]))
	dp[0] = triangle[0][0]
	for row := 1; row < rowNum; row++ {
		dp[row] = dp[row-1] + triangle[row][row]
		for col := row - 1; col > 0; col-- {
			dp[col] = min(dp[col-1], dp[col]) + triangle[row][col]
		}
		dp[0] = dp[0] + triangle[row][0]
	}
	ans := math.MaxInt
	for _, v := range dp {
		if v < ans {
			ans = v
		}
	}
	return ans
}
