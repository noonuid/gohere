package problem0172

func trailingZeroes(n int) int {
	ans := 0
	for multiple := 5; multiple <= n; multiple++ {
		for factor := multiple; factor%5 == 0; factor /= 5 {
			ans++
		}
	}
	return ans
}

func trailingZeroes_optimization(n int) int {
	ans := 0
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}
