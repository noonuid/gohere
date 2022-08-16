package operator_and_expression

import "fmt"

func Example_bitwise() {
	fmt.Println("0011 AND 0101 is", 0b0011&0b0101)
	fmt.Println("0011 OR 0101 is", 0b0011|0b0101)
	fmt.Println("0011 XOR 0101 is", 0b0011^0b0101)
	fmt.Println("1 << 5 is", 1<<5)
	fmt.Println("0x80 >> 2 is", 0x80>>2)

	// Output:
	// 0011 AND 0101 is 1
	// 0011 OR 0101 is 7
	// 0011 XOR 0101 is 6
	// 1 << 5 is 32
	// 0x80 >> 2 is 32
}
