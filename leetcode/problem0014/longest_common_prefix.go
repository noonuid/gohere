package problem0014

import "sort"

// 14. 最长公共前缀

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 先将 strs 排序，再比较头尾两个元素的最长公共前缀。
func longestCommonPrefix_sort(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	if strsLen == 1 {
		return strs[0]
	}
	sort.Strings(strs)
	head, tail := strs[0], strs[strsLen-1]
	i := 0
	for i < len(head) && i < len(tail) && head[i] == tail[i] {
		i++
	}
	return head[:i]
}

// 按列扫描。
func longestCommonPrefix_column_scan(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	colLen := len(strs[0])
	for col := 0; col < colLen; col++ {
		for row := 1; row < strsLen; row++ {
			if col == len(strs[row]) || strs[0][col] != strs[row][col] {
				return strs[0][:col]
			}
		}
	}
	return strs[0]
}
