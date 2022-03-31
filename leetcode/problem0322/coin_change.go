package problem0322

func coinChange(coins []int, amount int) int {
	var f = make([]int, amount+1)
	// 硬币个数不可能超过这个数
	var max = amount + 1

	var number = max
	for i := 1; i <= amount; i++ {
		number = max
		for _, coin := range coins {
			if i-coin >= 0 {
				number = min(number, f[i-coin]+1)
			}
		}
		f[i] = number
	}

	if f[amount] > amount {
		return -1
	} else {
		return f[amount]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
