package problem0096

// 96. 不同的二叉搜索树

// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

// 动态规划。
func numTrees(n int) int {
	// dp[i]：长度为 i 的序列能组成的不同二叉搜索树的个数。
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for m := 2; m <= n; m++ {
		for i := 1; i <= m; i++ {
			// dp[i-1] + dp[m-i]：以 i 为根、长度为 m 的序列组成的不同二叉搜索树的个数。
			dp[m] += dp[i-1] * dp[m-i]
		}
	}
	return dp[n]
}
