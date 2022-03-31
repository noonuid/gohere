package hj009

import "fmt"

func Do() {
	var n int
	var scanN, _ = fmt.Scan(&n)
	for scanN > 0 {
		fmt.Println(extract(n))

		scanN, _ = fmt.Scan(&n)
	}
}

func extract(n int) int {
	var bits = make([]int, 0)
	// 记录某一位是否已经存在
	var log = make(map[int]bool, n)
	for n > 0 {
		var bit = n % 10
		if _, ok := log[bit]; !ok {
			bits = append(bits, bit)
			log[bit] = true
		}
		n /= 10
	}

	// 根据每一位数字还原整数
	var result = 0
	for i := 0; i <= len(bits)-1; i++ {
		result = result*10 + bits[i]
	}

	return result
}
