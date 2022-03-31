package nc103

/**
 * 反转字符串
 * @param str string字符串
 * @return string字符串
 */
func solve(str string) string {
	var s = []byte(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
