package nc057

/**
 *
 * @param x int整型
 * @return int整型
 */
func reverse(x int) int {
	// write code here
	var result = 0
	for x != 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result > (1<<31)-1 || result < -(1<<31) {
		return 0
	}
	return result
}
