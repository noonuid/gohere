package concurrency

import (
	"fmt"
)

func Example_channel() {
	sum := func(s []int, c chan int) {
		sum := 0
		for _, v := range s {
			sum += v
		}
		c <- sum // 将和送入 c
	}

	{
		s := []int{7, 2, 8, -9, 4, 0}

		c := make(chan int)
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
		x, y := <-c, <-c // 从 c 中接收

		fmt.Println(x, y, x+y)
	}
	// Output:
	// -5 17 12
}

func Example_buffered_channel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// Output:
	// 1
	// 2
}

func Example_close_channel() {
	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}

	{
		c := make(chan int, 10)
		go fibonacci(cap(c), c)
		for i := range c {
			fmt.Println(i)
		}
	}
	// Output:
	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
}
