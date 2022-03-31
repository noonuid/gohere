package problem0438

func findAnagrams(s string, p string) []int {
	// 最终结果
	result := []int{}
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	// 记录字符串窗口和字符串 p 中各个字符出现的频率
	freq, freqP := make([]int, 26), make([]int, 26)

	for _, ch := range p {
		freqP[ch-'a']++
	}

	for slow, fast := 0, 0; fast < len(s); fast++ {
		freq[s[fast]-'a']++

		// 窗口中出现了某字符，但是p中没有
		// 窗口中某字符出现的次数大于 p 中字符出现的次数
		for freq[s[fast]-'a'] > freqP[s[fast]-'a'] {
			freq[s[slow]-'a']--
			slow++
		}

		if fast-slow+1 == len(p) {
			result = append(result, slow)
		}
	}

	return result
}
