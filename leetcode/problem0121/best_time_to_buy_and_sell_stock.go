package problem0121

import "math"

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

// 动态规划。
func maxProfit(prices []int) int {
	n, historyMin, profit := len(prices), math.MaxInt, 0
	for i := 0; i < n; i++ {
		if prices[i] < historyMin {
			historyMin = prices[i]
		} else if cur := prices[i] - historyMin; cur > profit {
			profit = cur
		}
	}
	return profit
}

// 暴力求解。
// 超出时间限制。
func maxProfit_brute_force(prices []int) int {
	n, profit := len(prices), 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if cur := prices[j] - prices[i]; profit < cur {
				profit = cur
			}
		}
	}
	return profit
}
