package problem0003

import "strings"

// 滑动窗口解法，使用 string 存储求解过程中的极大子串
func lengthOfLongestSubstring(s string) int {
	var substring = ""
	var maxLength = 0
	for _, ch := range s {
		if index := strings.IndexRune(substring, ch); index != -1 {
			// 此时的无重复字符子串长度是一个极大值，若大于 maxLength，则更新
			if len(substring) > maxLength {
				maxLength = len(substring)
			}
			// 缩短子串的左边界，排除重复字符
			substring = substring[index+1:]
		}
		// 延长子串的右边界
		substring += string(ch)
	}
	if maxLength > len(substring) {
		return maxLength
	} else {
		return len(substring)
	}
}

// 滑动窗口解法，使用 map 存储求解过程中的极大子串
func lengthOfLongestSubstringMap(s string) int {
	var appeared = make(map[byte]bool)
	var maxLength = 0

	// 子串的左右边界，right 不包含在子串中，代表即将加入到子串中的字符
	var left, right = 0, 0
	for right < len(s) {
		// 延长子串的右边界，加入新的不重复字符
		for right < len(s) && !appeared[s[right]] {
			appeared[s[right]] = true
			right++
		}

		// 此时的无重复字符子串长度是一个极大值，若大于 maxLength，则更新
		if right-left > maxLength {
			maxLength = right - left
		}

		// 缩短子串的左边界，排除重复字符
		for right < len(s) && appeared[s[right]] {
			delete(appeared, s[left])
			left++
		}
	}

	return maxLength
}

// 暴力求解
func lengthOfLongestSubstringForce(s string) int {
	// 该函数判断字符串中是否存在重复字符
	var hasRepeatingChar = func(s string) bool {
		for i := 0; i < len(s); i++ {
			for j := i + 1; j < len(s); j++ {
				if s[i] == s[j] {
					return true
				}
			}
		}
		return false
	}

	var maxLength = 0
	for i := 0; i < len(s); i++ {
		// 新的一轮迭代子串长度直接从 maxLength + 1 开始
		for j := i + maxLength + 1; j < len(s)+1; j++ {
			if !hasRepeatingChar(s[i:j]) {
				maxLength = j - i
			} else {
				break
			}
		}
	}
	return maxLength
}
