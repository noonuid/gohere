package problem0076

import (
	"math"
)

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

// 注意：

// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。

// 滑动窗口。
func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)
	if sLen < tLen {
		return ""
	}
	// tFs 记录 t 中字符出现的频率。
	tFs := make(map[byte]int)
	for i := 0; i < tLen; i++ {
		tFs[t[i]]++
	}
	// windowFs 记录 window 中 t 中字符出现的频率。
	windowFs := make(map[byte]int)
	isValid := func() bool {
		for ch, f := range tFs {
			if windowFs[ch] < f {
				return false
			}
		}
		return true
	}
	// findNext 找到 s 中下一个 t 中字符的位置。
	findNext := func(cur int) int {
		for ok := false; cur < sLen; cur++ {
			if _, ok = tFs[s[cur]]; ok {
				break
			}
		}
		return cur
	}
	left, right := 0, math.MaxInt
	// 找到 s 中第一个 t 中字符的位置。
	i := findNext(0)
	for j := i; j < sLen; j = findNext(j + 1) {
		windowFs[s[j]]++
		for j-i+1 >= tLen && isValid() {
			if j-i < right-left {
				left, right = i, j
			}
			windowFs[s[i]]--
			i = findNext(i + 1)
		}
	}
	if right == math.MaxInt {
		return ""
	}
	return s[left : right+1]
}

// 暴力枚举。超出时间限制.
func minWindow_brute_force(s string, t string) string {
	// tFs 记录 t 中字符出现得频率。
	tFs := make(map[rune]int)
	for _, ch := range t {
		tFs[ch]++
	}
	// isValid 判断 subStr 中是否包含 t 中所有字符。
	isValid := func(subStr string) bool {
		for ch, f := range tFs {
			count := 0
			for _, v := range subStr {
				if v == ch {
					count++
				}
			}
			if count < f {
				return false
			}
		}
		return true
	}
	left, right := 0, math.MaxInt
	for i := 0; i < len(s); i++ {
		for j := i + len(t); j <= len(s) && j-i < right-left; j++ {
			if isValid(s[i:j]) {
				left, right = i, j
			}
		}
	}
	if right == math.MaxInt {
		return ""
	} else {
		return s[left:right]
	}
}
