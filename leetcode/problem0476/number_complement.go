package problem0476

func findComplement(num int) int {
	temp := 1
	for temp <= num {
		temp <<= 1
	}
	return (temp - 1) ^ num
}
