package operator_and_expression

import "fmt"

func Example_logical() {
	fmt.Println("true AND false is", true && false)
	fmt.Println("true OR false is", true || false)
	fmt.Println("NOT true is", !true)

	// Output:
	// true AND false is false
	// true OR false is true
	// NOT true is false
}
