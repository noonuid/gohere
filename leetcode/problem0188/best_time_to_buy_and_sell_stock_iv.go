package problem0188

func maxProfit(k int, prices []int) int {
	buy, sell := make([]int, k), make([]int, k)
	for j := 0; j < k; j++ {
		buy[j] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		buy[0] = max(buy[0], -prices[i])
		sell[0] = max(sell[0], buy[0]+prices[i])
		for j := 1; j < len(buy); j++ {
			buy[j] = max(buy[j], sell[j-1]-prices[i])
			sell[j] = max(sell[j], buy[j]+prices[i])
		}
	}
	return sell[k-1]
}
