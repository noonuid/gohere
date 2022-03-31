package problem0005

// 暴力求解，枚举所有子串并判断是否为回文串
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	// 该函数判断 s 的指定区间子串是否为回文串
	var isPalindrome = func(begin, end int) bool {
		if begin == end {
			return true
		}

		for i, j := begin, end-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	var begin, end = 0, 0
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			if (end-begin) < (j-i) && isPalindrome(i, j) {
				begin = i
				end = j
				// 内层循环为倒序遍历，后序轮次中找到的回文子串只会更短
				break
			}
		}
	}

	return s[begin:end]
}

// 动态规划
func longestPalindromeDynamicProgramming(s string) string {
	var length = len(s)
	if length <= 1 {
		return s
	}

	// 状态矩阵，f[i][j] 表示 s[i: j+1] 是否是回文串
	var f = make([][]bool, length)
	for i := 0; i < length; i++ {
		f[i] = make([]bool, length)
		// 单个字符一定是回文串
		// 此行也可省略，因为后面的边界长度为 2，不会参考长度为 1 的状态
		f[i][i] = true
	}

	var begin, end = 0, 0
	for subLength := 2; subLength <= length; subLength++ {
		for i := 0; i+subLength-1 < length; i++ {
			var j = i + subLength - 1

			if s[i] == s[j] {
				if subLength <= 3 {
					f[i][j] = true
				} else {
					f[i][j] = f[i+1][j-1]
				}

				// 如果当前子串是回文串并且长度大于已记录的最大长度，则更新
				if f[i][j] && subLength > end-begin+1 {
					begin = i
					end = j
				}
			} else {
				f[i][j] = false
			}
		}
	}

	return s[begin : end+1]
}

// 中心扩展算法
func longestPalindromeCenter(s string) string {
	if len(s) <= 1 {
		return s
	}

	// 该函数计算从回文串中心处最大可以扩展的长度，返回左右两个边界
	var expandAroundCenter = func(left, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		return left + 1, right - 1
	}

	var begin, end = 0, 0
	for i := 0; i < len(s); i++ {
		// 以回文串长度为奇数和偶数分别进行扩展
		var leftOdd, rightOdd = expandAroundCenter(i, i)
		var leftEven, rightEven = expandAroundCenter(i, i+1)

		if rightOdd-leftOdd+1 > end-begin+1 {
			begin, end = leftOdd, rightOdd
		}
		if rightEven-leftEven+1 > end-begin+1 {
			begin, end = leftEven, rightEven
		}
	}

	return s[begin : end+1]
}
