package recursion

// 使用递归方法计算一个数的阶乘
func Factorial(num int) int {
	if num < 0 {
		return 0
	}

	if num == 0 || num == 1 {
		return 1
	}

	return num * Factorial(num-1)
}

// 使用循环方法计算一个数的阶乘
func FactorialIterative(num int) int {
	if num < 0 {
		return 0
	}

	var factorial = 1
	for i := num; i > 1; i-- {
		factorial = factorial * i
	}

	return factorial
}

// 使用递归方法计算斐波那契数列
func Fibonacci(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}

// 使用缓存方法计算斐波那契数列
func FibonacciMemorization(n int) int {
	if n <= 0 {
		return 0
	}

	var memo = make([]int, n+1)
	memo[0] = 0
	memo[1] = 1

	return fibonacciMemorization(n, memo)
}

// 供 FibonacciMemorization 内部调用实现递归
func fibonacciMemorization(n int, memo []int) int {
	if n == 0 {
		return 0
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = fibonacciMemorization(n-2, memo) + fibonacciMemorization(n-1, memo)
	return memo[n]
}

// 使用循环方法计算斐波那契数列
func FibonacciIterative(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	var num1, num2, num3 = 1, 1, 2
	for i := 3; i <= n; i++ {
		num3 = num1 + num2
		num1 = num2
		num2 = num3
	}

	return num3
}
