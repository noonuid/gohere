package problem0412

import (
	"strconv"
	"strings"
)

// 412. Fizz Buzz

// 给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1 开始）返回结果，其中：

// answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
// answer[i] == "Fizz" 如果 i 是 3 的倍数。
// answer[i] == "Buzz" 如果 i 是 5 的倍数。
// answer[i] == i （以字符串形式）如果上述条件全不满足。

// 字符串拼接。
func fizzBuzz_string_concatenation(n int) []string {
	answer := make([]string, n)
	for i := 1; i < n+1; i++ {
		sb := &strings.Builder{}
		remainder3, remainder5 := i%3, i%5
		if remainder3 == 0 {
			sb.WriteString("Fizz")
		}
		if remainder5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		answer[i-1] = sb.String()
	}
	return answer
}

func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 1; i < n+1; i++ {
		remainder3, remainder5 := i%3, i%5
		switch {
		case remainder3 == 0 && remainder5 == 0:
			answer[i-1] = "FizzBuzz"
		case remainder3 == 0:
			answer[i-1] = "Fizz"
		case remainder5 == 0:
			answer[i-1] = "Buzz"
		default:
			answer[i-1] = strconv.Itoa(i)
		}
	}
	return answer
}
