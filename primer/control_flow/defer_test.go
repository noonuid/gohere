package control_flow

import "fmt"

func Example_defer() {
	defer fmt.Println("world")

	fmt.Println("hello")

	// Output:
	// hello
	// world
}

func Example_stacking_defers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// Output:
	// counting
	// done
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}
