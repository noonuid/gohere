package problem0005

// 给你一个字符串 s，找到 s 中最长的回文子串。
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

// 暴力求解，枚举所有子串并判断是否为回文串。
func longestPalindrome(s string) string {
	// 判断字符串是否是回文字符串。
	isPalindrome := func(s string) bool {
		left, right := 0, len(s)-1
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left, right = left+1, right-1
		}
		return true
	}

	begin, end := 0, 0
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j-i > end-begin; j-- {
			if isPalindrome(s[i : j+1]) {
				begin, end = i, j
				break
			}
		}
	}
	return s[begin : end+1]
}

// 动态规划。
func longestPalindromeDynamicProgramming(s string) string {
	length := len(s)
	// 状态矩阵，f[i][j] 表示 s[i: j+1] 是否是回文串
	f := make([][]bool, length)
	// 初始化状态矩阵。
	for i := 0; i < length; i++ {
		f[i] = make([]bool, length)
		// 单个字符一定是回文串。
		// 此行也可省略，因为当子串长度为 2 或 3 时，不会参考 f 的值。
		// f[i][i] = true
	}

	begin, end := 0, 0
	for subLen := 2; subLen <= length; subLen++ {
		for i := 0; i+subLen-1 < length; i++ {
			j := i + subLen - 1
			if s[i] == s[j] {
				if subLen <= 3 {
					f[i][j] = true
				} else {
					f[i][j] = f[i+1][j-1]
				}
				// 如果当前子串是回文串并且长度大于已记录的最大长度，则更新。
				if f[i][j] && subLen > end-begin+1 {
					begin, end = i, j
				}
			} else {
				f[i][j] = false
			}
		}
	}
	return s[begin : end+1]
}

// 中心扩展算法。
func longestPalindromeCenter(s string) string {
	begin, end := 0, 0
	for center := 0; center < len(s); center++ {
		// 回文字符串长度为奇数。
		left, right := center-1, center+1
		for left > -1 && right < len(s) && s[left] == s[right] {
			left, right = left-1, right+1
		}
		left, right = left+1, right-1
		if end-begin < right-left {
			begin, end = left, right
		}
		// 回文字符串长度为偶数。
		left, right = center-1, center
		for left > -1 && right < len(s) && s[left] == s[right] {
			left, right = left-1, right+1
		}
		left, right = left+1, right-1
		if end-begin < right-left {
			begin, end = left, right
		}
	}

	return s[begin : end+1]
}
