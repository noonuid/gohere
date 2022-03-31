package problem1758

func minOperations(s string) int {
	count1, count2 := 0, 0

	for index, ch := range s {
		if int(ch)%2 != index%2 {
			// 偶数位为0，奇数位为1
			count1++
		} else {
			// 偶数位为1，奇数位为0
			count2++
		}
	}

	if count1 < count2 {
		return count1
	} else {
		return count2
	}
}
