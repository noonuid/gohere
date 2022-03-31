package hj038

import "fmt"

func Do() {
	var initHeight = 0
	var n, _ = fmt.Scan(&initHeight)
	for n > 0 {
		var distance = -float64(initHeight)
		var lastHeight = float64(initHeight)
		for i := 0; i < 5; i++ {
			distance += lastHeight * 2
			lastHeight /= 2
		}
		fmt.Println(distance)
		fmt.Println(lastHeight)

		n, _ = fmt.Scan(&initHeight)
	}
}
