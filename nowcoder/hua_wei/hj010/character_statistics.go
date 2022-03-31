package hj010

import (
	"fmt"
)

// 方法一
func Do() {
	var input string
	var n, _ = fmt.Scan(&input)
	for n > 0 {
		var table = make(map[rune]int, 128)
		var freqency = 0
		for _, ch := range input {
			if _, ok := table[ch]; !ok {
				freqency += 1
			}
			table[ch] += 1
		}

		fmt.Println(freqency)

		n, _ = fmt.Scan(&input)
	}
}

// 方法二
func DoI() {
	var input string
	var n, _ = fmt.Scan(&input)
	for n > 0 {
		var table = make(map[rune]bool, 128)
		for _, ch := range input {
			table[ch] = true
		}
		fmt.Println(len(table))

		n, _ = fmt.Scan(&input)
	}
}
