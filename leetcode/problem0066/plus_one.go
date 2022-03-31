package problem0066

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++

		if digits[i] != 10 {
			return digits
		}

		digits[i] = 0
	}

	digits = append(digits, 0)
	digits[0] = 1
	return digits
}
