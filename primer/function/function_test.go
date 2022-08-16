package function

import "fmt"

func Example_function() {
	var add = func(x int, y int) int {
		return x + y
	}

	{
		fmt.Println(add(42, 13))
	}
	// Output:
	// 55
}

func Example_share_a_type() {
	var add = func(x, y int) int {
		return x + y
	}

	{
		fmt.Println(add(42, 13))
	}
	// Output:
	// 55
}

func Example_multiple_results() {
	var swap = func(x, y string) (string, string) {
		return y, x
	}

	{
		a, b := swap("hello", "world")
		fmt.Println(a, b)
	}
	// Output:
	// world hello
}

func Example_named_returned_values() {
	var split = func(sum int) (x, y int) {
		x = sum * 4 / 9
		y = sum - x
		return
	}

	{
		fmt.Println(split(17))
	}
	// Output:
	// 7 10
}
