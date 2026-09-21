package problem0392

func isSubsequence(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	if lenS == 0 {
		return true
	} else if lenT == 0 || lenS > lenT {
		return false
	}
	dp := make([][]bool, lenS)
	for i := 0; i < lenS; i++ {
		dp[i] = make([]bool, lenT)
	}
	if s[0] == t[0] {
		dp[0][0] = true
	}
	for j := 1; j < lenT; j++ {
		if dp[0][j-1] || s[0] == t[j] {
			dp[0][j] = true
		}
	}
	for i := 1; i < lenS; i++ {
		for j := 1; j < lenT; j++ {
			if dp[i][j-1] || (s[i] == t[j] && dp[i-1][j-1]) {
				dp[i][j] = true
			}
		}
	}
	return dp[lenS-1][lenT-1]
}

func isSubsequence_two_pointer(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	i, j := 0, 0
	for i < lenS && j < lenT {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == lenS
}

func isSubsequence_preprocess(s string, t string) bool {
	lenS, lenT := len(s), len(t)
	if lenS == 0 {
		return true
	} else if lenT == 0 || lenS > lenT {
		return false
	}
	dp := make([][26]int, lenT+1)
	for j := 0; j < 26; j++ {
		dp[lenT][j] = lenT
	}
	for i := lenT - 1; i > -1; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	tPos := 0
	for i := 0; i < lenS; i++ {
		if dp[tPos][s[i]-'a'] == lenT {
			return false
		}
		tPos = dp[tPos][s[i]-'a'] + 1
	}
	return true
}
