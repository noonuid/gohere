package hj060

import (
	"fmt"
)

func Do() {
	var even, num1, num2 int
	var number, _ = fmt.Scan(&even)
	for number > 0 {
		num1, num2 = 0, even
		for i := 2; i <= even/2; i++ {
			if isPrime(i) && isPrime(even-i) && even-i-i < num2-num1 {
				num1, num2 = i, even-i
			}
		}
		fmt.Println(num1)
		fmt.Println(num2)

		number, _ = fmt.Scan(&even)
	}
}

// 判断一个数是否为素数
func isPrime(x int) bool {
	if x == 1 || (x > 2 && x%2 == 0) {
		return false
	}

	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}
