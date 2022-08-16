package boolean

import "fmt"

func Example_bool() {
	fmt.Println("true AND false is {}", true && false)
	fmt.Println("true OR false is {}", true || false)
	fmt.Println("NOT true is {}", !true)
	// OUtput:
	// true AND false is {} false
	// true OR false is {} true
	// NOT true is {} false
}
