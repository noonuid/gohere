package problem1556

import "strconv"

func thousandSeparator(n int) string {
	var result = []byte{}
	var str = strconv.Itoa(n)
	var count = len(str) % 3
	if count == 0 {
		count = 3
	}
	for i := 0; i < len(str); i++ {
		result = append(result, str[i])
		count--
		if count == 0 && i != len(str)-1 {
			result = append(result, '.')
			count = 3
		}
	}

	return string(result)
}
