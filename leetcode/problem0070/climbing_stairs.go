package problem0070

// 动态规划
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var f = make([]int, n+1)
	f[1], f[2] = 1, 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-2] + f[i-1]
	}

	return f[n]
}

// 动态规划，不使用状态数组 f[i] 存储过程值
func climbStairsScroll(n int) int {
	if n <= 2 {
		return n
	}

	var a, b, c = 1, 2, 3
	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}
