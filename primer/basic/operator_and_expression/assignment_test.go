package operator_and_expression

import "fmt"

func Example_assignment() {
	var x = 0
	fmt.Println("The value of x is:", x)
	x = 1
	fmt.Println("The value of x is:", x)

	// Output:
	// 	The value of x is: 0
	// The value of x is: 1
}
