package nc126

/**
 * 最少货币数
 * @param arr int整型一维数组 the array
 * @param aim int整型 the target
 * @return int整型
 */
func minMoney(arr []int, aim int) int {
	// write code here
	var f = make([]int, aim+1)
	var max = aim + 1

	var number = max
	for i := 1; i <= aim; i++ {
		number = max
		for _, coin := range arr {
			if i-coin >= 0 {
				number = min(number, f[i-coin]+1)
			}
		}
		f[i] = number
	}

	if f[aim] > aim {
		return -1
	} else {
		return f[aim]
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
