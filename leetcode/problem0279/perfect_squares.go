package problem0279

// 279. 完全平方数

// 给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

func numSquares(n int) int {
	f := make([]int, n+1)
	for num := 1; num < n+1; num++ {
		min := num
		for i := 1; i*i <= num; i++ {
			if min > 1+f[num-i*i] {
				min = 1 + f[num-i*i]
			}
		}
		f[num] = min
	}
	return f[n]
}
