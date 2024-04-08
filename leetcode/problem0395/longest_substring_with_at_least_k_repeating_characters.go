package problem0395

import "strings"

// 395. 至少有 K 个重复字符的最长子串

// 给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。

// 如果不存在这样的子字符串，则返回 0。

// 枚举。
// 超出时间限制。
func longestSubstring_enum(s string, k int) int {
	n := len(s)
	max := 0
	isValid := func(sub string) bool {
		fs := make([]int, 26)
		for _, ch := range sub {
			fs[ch-'a']++
		}
		for _, f := range fs {
			if 0 < f && f < k {
				return false
			}
		}
		return true
	}
	for left := 0; left < n; left++ {
		for right := n; right-left > max; right-- {
			if isValid(s[left:right]) && max < len(s[left:right]) {
				max = len(s[left:right])
			}
		}
	}
	return max
}

// 分治法。
func longestSubstring(s string, k int) int {
	fs := make([]int, 26)
	for _, ch := range s {
		fs[ch-'a']++
	}
	var sep rune
	for i, f := range fs {
		if 0 < f && f < k {
			sep = 'a' + rune(i)
			break
		}
	}
	if sep == 0 {
		return len(s)
	}
	max := 0
	substrs := strings.Split(s, string(sep))
	for _, substr := range substrs {
		if max < len(substr) {
			if cur := longestSubstring(substr, k); max < cur {
				max = cur
			}
		}
	}
	return max
}
