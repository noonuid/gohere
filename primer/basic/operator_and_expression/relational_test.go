package operator_and_expression

import "fmt"

func Example_relational() {
	fmt.Println("0 == 1 is", 0 == 1)
	fmt.Println("0 != 1 is", 0 != 1)
	fmt.Println("0 > 1 is", 0 > 1)
	fmt.Println("0 < 1 is", 0 < 1)
	fmt.Println("0 <= 1 is", 0 <= 1)
	fmt.Println("0 >= 1 is", 0 >= 1)

	// Output:
	// 0 == 1 is false
	// 0 != 1 is true
	// 0 > 1 is false
	// 0 < 1 is true
	// 0 <= 1 is true
	// 0 >= 1 is false
}
