package hj008

import (
	"fmt"
	"sort"
)

func Do() {
	var n = 0
	// 读取记录的总数量
	fmt.Scan(&n)

	var records = make(map[int]int, n)
	var index, value = 0, 0
	// 读取所有记录行
	for n > 0 {
		fmt.Scan(&index, &value)
		records[index] += value

		n--
	}

	// 利用切片对记录的索引排序
	var indexes = make([]int, len(records))
	var i = 0
	for key := range records {
		indexes[i] = key
		i++
	}
	sort.Ints(indexes)

	// 输出结果
	for j := 0; j < i; j++ {
		fmt.Println(indexes[j], records[indexes[j]])
	}
}
