package hj108

import "fmt"

func Do() {
	var a, b = 0, 0

	fmt.Scan(&a, &b)
	fmt.Println(findTheLeastCommonMultiple(a, b))
}

func findTheLeastCommonMultiple(a, b int) int {
	var bigger = max(a, b)

	for i := bigger; i < a*b; i++ {
		if i%a == 0 && i%b == 0 {
			return i
		}
	}

	return a * b
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
