package problem1143

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	f := make([][]int, len1+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, len2+1)
	}
	for i := 1; i < len(f); i++ {
		for j := 1; j < len(f[0]); j++ {
			if text1[i-1] == text2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return f[len1][len2]
}
