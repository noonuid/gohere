package flow_control

import "fmt"

func Example_for() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Output:
	// 45
}

func Example_optional_statements() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Output:
	// 1024
}

func Example_for_is_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Output:
	// 1024
}
