package string

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串定义
	var str = "abc"

	// 字符串长度
	fmt.Println("字符串长度：", len(str))
	fmt.Println()

	// 字符串拼接
	str = str + "abc" + "abc" + "abc"
	fmt.Println("字符串拼接：", str)
	fmt.Println()

	// 字符串替换
	fmt.Println("字符串替换（前 n 个）：", strings.Replace(str, "abc", "xyz", 2))
	fmt.Println("字符串替换（全部）：", strings.Replace(str, "abc", "xyz", -1))
	fmt.Println("字符串替换（全部）：", strings.ReplaceAll(str, "abc", "xyz"))
	fmt.Println()

	// 字符串查找
	fmt.Println("字符串查找（是否包含字串）：", strings.Contains(str, "abc"))
	fmt.Println("字符串查找（是否包含字串中的任意字符）：", strings.ContainsAny(str, "ayz"))
	fmt.Println("字符串查找（返回字符位置）：", strings.IndexRune(str, 'b'))
	fmt.Println()

	// 字符串截取
	fmt.Println("字符串截取：", str[:3])
	fmt.Println()

	// 字符串统计
	fmt.Println("字符串统计：", strings.Count(str, "abc"))
	fmt.Println()

	// 字符串的拆分与合并
	fmt.Println("字符串的拆分：", strings.Split(str, "b"))
	fmt.Println("字符串的拆分：", strings.Fields("abc abc abc abc"))
	fmt.Println("字符串的合并：", strings.Join([]string{"ab", "ab", "ab", "ab"}, "c"))
	fmt.Println()

	// 字符串转大写
	str = strings.ToUpper(str)
	fmt.Println("字符串转大写：", str)
	fmt.Println()

	// 字符串转小写
	str = strings.ToLower(str)
	fmt.Println("字符串转小写：", str)
	fmt.Println()

	// 字符串比较
	fmt.Println(strings.Compare("abc", "abc"))
	fmt.Println(strings.Compare("abc", "aBc"))
	fmt.Println(strings.Compare("aBc", "abc"))
}
