package problem0017

// 17. 电话号码的字母组合

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

// 回溯法。
func letterCombinations(digits string) []string {
	n := len(digits)
	res := []string{}
	if n == 0 {
		return res
	}
	numMap := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var backtrack func(path []byte, index int)
	backtrack = func(path []byte, index int) {
		if len(path) == n {
			res = append(res, string(path))
			return
		}
		letters := numMap[digits[index]]
		ltsLen := len(letters)
		for i := 0; i < ltsLen; i++ {
			path = append(path, letters[i])
			backtrack(path, index+1)
			path = path[:len(path)-1]
		}
	}

	backtrack([]byte{}, 0)
	return res
}
