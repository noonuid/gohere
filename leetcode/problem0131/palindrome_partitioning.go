package problem0131

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。返回 s 所有可能的分割方案。

// 回溯法。
func partition(s string) [][]string {
	length, result := len(s), [][]string{}
	isPalindrome := func(left, right int) bool {
		for i, j := left, right; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	var backtrack func(path []string, left int)
	backtrack = func(path []string, left int) {
		if left >= length {
			result = append(result, append([]string{}, path...))
			return
		}
		for right := left; right < length; right++ {
			if left == right || isPalindrome(left, right) {
				backtrack(append(path, s[left:right+1]), right+1)
			}
		}
	}
	backtrack([]string{}, 0)
	return result
}

// 回溯法+动态规划预处理。
func partition_backtrack_dp(s string) [][]string {
	length, result := len(s), [][]string{}
	f := make([][]bool, length)
	for i := 0; i < length; i++ {
		f[i] = make([]bool, length)
		for j := 0; j < length; j++ {
			f[i][j] = true
		}
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	var backtrack func(path []string, left int)
	backtrack = func(path []string, left int) {
		if left >= length {
			result = append(result, append([]string{}, path...))
			return
		}
		for right := left; right < length; right++ {
			if f[left][right] {
				backtrack(append(path, s[left:right+1]), right+1)
			}
		}
	}
	backtrack([]string{}, 0)
	return result
}
