package problem0461

import "math/bits"

// 461. 汉明距离

// 两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

// 给你两个整数 x 和 y，计算并返回它们之间的汉明距离。

// 移位计数法。
func hammingDistance_shift_and_count(x int, y int) int {
	count := 0
	for xor := x ^ y; xor > 0; xor >>= 1 {
		count += xor & 1
	}
	return count
}

// 除二取余法。
func hammingDistance_divide_by_two(x int, y int) int {
	dis := 0
	xor := x ^ y
	for xor > 0 {
		if xor%2 == 1 {
			dis++
		}
		xor /= 2
	}
	return dis
}

// OnesCount。
func hammingDistance_ones_count(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
