package problem0139

// 139. 单词拆分

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

// 动态规划。
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i < n+1; i++ {
		for _, word := range wordDict {
			if preLen := i - len(word); preLen >= 0 && f[preLen] && s[preLen:i] == word {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}
