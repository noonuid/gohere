package problem0049

import "sort"

// 49. 字母异位词分组

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

// 使用 map 存储同一组字母异位词。
func groupAnagrams_map(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
		sorted := string(bytes)
		m[sorted] = append(m[sorted], str)
	}
	res := make([][]string, 0, len(m))
	for _, anagrams := range m {
		res = append(res, anagrams)
	}
	return res
}

// 暴力枚举。
func groupAnagrams_brute(strs []string) [][]string {
	length := len(strs)
	res := [][]string{}
	sorted := make([]string, length)
	for i := 0; i < length; i++ {
		bytes := []byte(strs[i])
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
		sorted[i] = string(bytes)
	}
	used := make([]bool, length)
	for i := 0; i < length; i++ {
		if used[i] {
			continue
		}
		anagrams := []string{strs[i]}
		used[i] = true
		for j := i + 1; j < length; j++ {
			if sorted[i] == sorted[j] {
				anagrams = append(anagrams, strs[j])
				used[j] = true
			}
		}
		res = append(res, anagrams)
	}
	return res
}
