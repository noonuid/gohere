package hj003

import (
	"fmt"
	"sort"
)

func Do() {
	var n int
	var scanN, _ = fmt.Scan(&n)
	for scanN > 0 {
		var numbers = make([]int, n)
		for n > 0 {
			fmt.Scan(&numbers[len(numbers)-n])

			n--
		}

		sort.Ints(numbers)
		fmt.Println(numbers[0])
		for i := 1; i < len(numbers); i++ {
			if numbers[i] > numbers[i-1] {
				fmt.Println(numbers[i])
			}
		}

		scanN, _ = fmt.Scan(&n)
	}
}
