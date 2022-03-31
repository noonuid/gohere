package nc055

/**
 *
 * @param strs string字符串一维数组
 * @return string字符串
 */
func longestCommonPrefix(strs []string) string {
	// write code here
	if len(strs) == 0 {
		return ""
	}

	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
	}

	return prefix
}

func commonPrefix(str1, str2 string) string {
	var i = 0
	for i < len(str1) && i < len(str2) {
		if str1[i] != str2[i] {
			break
		} else {
			i++
		}
	}
	return str1[:i]
}
