package problem0309

// 309. 买卖股票的最佳时机含冷冻期

// 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​

// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 动态规划。
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// f[i][0]：第 i 天结束以后，手中持有股票对应的最大收益。
	// f[i][1]：第 i 天结束以后，手中没有股票，并且处于冷冻期对应的最大收益。
	// f[i][2]：第 i 天结束以后，手中没有股票，并且没有处于冷冻期对应的最大收益。
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i])
		f[i][1] = f[i-1][0] + prices[i]
		f[i][2] = max(f[i-1][1], f[i-1][2])
	}
	return max(f[n-1][1], f[n-1][2])
}
