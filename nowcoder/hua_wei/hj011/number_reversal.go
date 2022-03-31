package hj011

import (
	"bufio"
	"fmt"
	"os"
)

func Do() {
	var reader = bufio.NewReader(os.Stdin)
	var input, _, _ = reader.ReadLine()
	for len(input) != 0 {
		for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
			input[i], input[j] = input[j], input[i]
		}
		fmt.Println(string(input))

		input, _, _ = reader.ReadLine()
	}
}
