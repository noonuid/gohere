package array

import "fmt"

// 数组声明与初始化
func Declare() {
	var strs [2]string
	strs[0] = "Hello"
	strs[1] = "World"
	// [Hello World]
	fmt.Println(strs)

	var primes = [6]int{2, 3, 5, 7, 11, 13}
	// [2 3 5 7 11 13]
	fmt.Println(primes)
}
