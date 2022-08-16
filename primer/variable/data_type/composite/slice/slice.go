package slice

import (
	"fmt"
	"sort"
)

func main() {
	// 切片定义
	var s = []int{1, 3, 5, 7, 9, 11}
	fmt.Println("切片定义：", s)
	fmt.Println()

	// 切片长度与容量
	fmt.Println("切片长度与容量")
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println()

	// 修改切片对数组的影响
	names := [4]string{"John", "Paul", "George", "Ringo"}
	a, b := names[0:2], names[1:3]
	b[0] = "XXX"
	fmt.Println("修改切片对数组的影响")
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println()

	// 数组排序
	s = []int{5, 2, 6, 3, 1, 4}
	fmt.Println("数组排序")
	sort.Ints(s)
	fmt.Println(s)
	s = []int{5, 2, 6, 3, 1, 4}
	sort.Slice(s, func(i int, j int) bool { return s[i] > s[j] })
	fmt.Println(s)
	fmt.Println()

	// 数组查找
	s = []int{1, 3, 5, 7, 9, 11}
	fmt.Println("数组查找")
	fmt.Println(sort.SearchInts(s, 5))
	fmt.Println(sort.Search(len(s), func(i int) bool { return s[i] >= 5 }))
}
