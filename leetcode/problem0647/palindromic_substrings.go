package problem0647

// 647. 回文子串

// 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。

// 回文字符串 是正着读和倒过来读一样的字符串。

// 子字符串 是字符串中的由连续字符组成的一个序列。

// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

// 中心扩展。
func countSubstrings_center_expand(s string) int {
	count, length := 0, len(s)
	for i := 0; i < 2*length-1; i++ {
		left, right := i/2, i/2+i%2
		for left > -1 && right < length && s[left] == s[right] {
			count++
			left, right = left-1, right+1
		}
	}
	return count
}

// 动态规划。
func countSubstrings_dynamic_programming(s string) int {
	count, length := len(s), len(s)
	f := make([][]bool, length)
	for i := 0; i < length; i++ {
		f[i] = make([]bool, length)
		f[i][i] = true
	}
	// 从下到上，从左到右遍历。
	for row := length - 1; row > -1; row-- {
		for col := row + 1; col < length; col++ {
			if s[row] == s[col] {
				if col-row == 1 || f[row+1][col-1] {
					f[row][col] = true
					count++
				}
			}
		}
	}
	return count
}

// 暴力枚举。
func countSubstrings_brute(s string) int {
	count, length := len(s), len(s)

	valid := func(start, end int) bool {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	for start := 0; start < length-1; start++ {
		for end := start + 1; end < length; end++ {
			if valid(start, end) {
				count++
			}
		}
	}
	return count
}
