package problem0097

// 97. 交错字符串

// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空
// 子字符串
// ：

// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
// 注意：a + b 意味着字符串 a 和 b 连接。

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	f := make([][]bool, m+1)
	f[0] = make([]bool, n+1)
	f[0][0] = true
	for i := 1; i < m+1; i++ {
		f[i] = make([]bool, n+1)
		f[i][0] = f[i-1][0] && (s1[i-1] == s3[i-1])
	}
	for j := 1; j < n+1; j++ {
		f[0][j] = f[0][j-1] && (s2[j-1] == s3[j-1])
		if !f[0][j] {
			break
		}
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			f[i][j] = (f[i-1][j] && s1[i-1] == s3[i+j-1]) || (f[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return f[m][n]
}
