package basic

import "fmt"

func Example_variable() {
	var c, python, java bool

	{
		var i int
		fmt.Println(i, c, python, java)
	}
	// OUtput:
	// 0 false false false
}

func Example_initialize_variable() {
	var i, j int = 1, 2

	{
		var c, python, java = true, false, "no!"
		fmt.Println(i, j, c, python, java)
	}
	// OUtput:
	// 1 2 true false no!
}

func Example_short_variable_declaration() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	// OUtput:
	// 1 2 3 true false no!
}
