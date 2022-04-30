package data_type

import "fmt"

func Example_variable() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// OUtput:
	// 0 0 false ""
}
