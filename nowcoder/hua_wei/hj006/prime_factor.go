package hj006

import "fmt"

func Do() {
	var n int
	var number, _ = fmt.Scan(&n)
	for number > 0 {
		PrimeFactor(n)

		number, _ = fmt.Scan(&n)
	}
}

// 导出该函数是为了使用 Example 测试
func PrimeFactor(n int) {
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			if i != n {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("%d\n", i)
			}
			n /= i
		}
	}
	// 如果 m=1，则循环中刚好完成质数分解，如果 m>1，说明没有完成分解，m 就是最后一个质数
	//同时，这句也可以应对输入为质数的特殊情况
	if n >= 2 {
		fmt.Printf("%d\n", n)
	}
}
