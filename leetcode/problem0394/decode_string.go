package problem0394

import (
	"strconv"
	"strings"
)

// 394. 字符串解码

// 给定一个经过编码的字符串，返回它解码后的字符串。

// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

// 使用栈操作。
func decodeString_stack(s string) string {
	n := len(s)
	str := ""
	numStack, strStack := []int{}, []string{}
	for i := 0; i < n; i++ {
		switch {
		case 'a' <= s[i] && s[i] <= 'z':
			// 记录左括号（或者 s 开头）与数字之间的子字符串。
			str += string(s[i])
		case '0' <= s[i] && s[i] <= '9':
			// 将数字前的子字符串入栈以供之后使用。
			strStack = append(strStack, str)
			str = ""
			// 提取出次数并将次数入栈。
			numLeft := i
			for i = i + 1; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
			}
			num, _ := strconv.Atoi(s[numLeft:i])
			numStack = append(numStack, num)
		case s[i] == ']':
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str = strings.Repeat(str, num)
			// 重复后的子字符串需要加上之前由对应的【重复次数数字】触发入栈的子字符串。
			str = strStack[len(strStack)-1] + str
			strStack = strStack[:len(strStack)-1]
		}
	}
	return str
}
