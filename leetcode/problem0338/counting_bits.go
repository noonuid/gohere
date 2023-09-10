package problem0338

// 338. 比特位计数

// 给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。

// 动态规划。
func countBits_dynamic_programming(n int) []int {
	f := make([]int, n+1)
	highBit := 0
	for i := 1; i < n+1; i++ {
		// i 的二进制表示只有最高位为 1，其余全为 0。
		if i&(i-1) == 0 {
			// 更新最高有效位。
			highBit = i
		}
		// i 的二进制表示只比 i-highBit 多一位 1。
		f[i] = f[i-highBit] + 1
	}
	return f
}

// Brian Kernighan 算法。
func countBits_brian_kernighan(n int) []int {
	countOnes := func(x int) int {
		res := 0
		for x > 0 {
			// 对于任意整数 x，令 x=x&(x−1)，该运算将 x 的二进制表示的最后一个 1 变成 0。
			x = x & (x - 1)
			res++
		}
		return res
	}
	res := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		res[i] = countOnes(i)
	}
	return res
}
