package conversion

import (
	"fmt"
	"strconv"
)

func main() {
	// string <-> byte
	fmt.Println("string <-> byte")
	fmt.Printf("string -> byte: %c\n", byte("abc"[0]))
	fmt.Printf("byte -> string: %s\n", string('a'))
	fmt.Println()

	// string <-> int
	fmt.Println("string <-> int")
	fmt.Println(strconv.Atoi("65"))
	fmt.Println(strconv.ParseInt("65", 10, 0))
	fmt.Println(strconv.Itoa(65))
	fmt.Println(strconv.FormatInt(65, 10))
	fmt.Println()

	// byte <-> int
	fmt.Println("byte <-> int")
	fmt.Println("byte -> ASCII: ", int('0'))
	fmt.Println("ASCII -> byte: ", byte(48))
	fmt.Println("byte -> int: ", int('1'-'0'))
	fmt.Println("int -> byte: ", byte(1+'0'))
	fmt.Println()

	// string <-> []byte
	fmt.Println("string <-> []byte")
	fmt.Println([]byte("abc"))
	fmt.Println(string([]byte{'a', 'b', 'c'}))
	fmt.Println()
}
