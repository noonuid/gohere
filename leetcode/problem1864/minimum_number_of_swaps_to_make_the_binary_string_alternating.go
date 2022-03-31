package problem1864

import (
	"math"
	"strings"
)

func minSwaps(s string) int {
	n := len(s)
	n0 := strings.Count(s, "0")
	n1 := strings.Count(s, "1")

	// 无法完成转化
	if n0-n1 > 1 || n1-n0 > 1 {
		return -1
	}

	result := math.MaxUint16

	// 转化为 "0101..."
	if n0 == (n+1)/2 && n1 == n/2 {
		count0 := 0
		for index, ch := range s {
			if int(ch-'0') != (index % 2) {
				count0++
			}
		}

		if count0/2 < result {
			result = count0 / 2
		}
	}
	// 转化为 "1010..."
	if n1 == (n+1)/2 && n0 == n/2 {
		count1 := 0
		for index, ch := range s {
			if int(ch-'0') == (index % 2) {
				count1++
			}
		}

		if count1/2 < result {
			result = count1 / 2
		}
	}

	return result
}
