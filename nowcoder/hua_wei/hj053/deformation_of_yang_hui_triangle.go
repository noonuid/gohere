package hj053

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Do() {
	var reader = bufio.NewReader(os.Stdin)
	var input, _, _ = reader.ReadLine()
	for len(input) != 0 {
		var n, _ = strconv.Atoi(string(input))
		if n == 1 || n == 2 {
			fmt.Println(-1)
		} else if (n-2)%4 == 1 {
			fmt.Println(2)
		} else if (n-2)%4 == 2 {
			fmt.Println(3)
		} else if (n-2)%4 == 3 {
			fmt.Println(2)
		} else if (n-2)%4 == 0 {
			fmt.Println(4)
		}

		input, _, _ = reader.ReadLine()
	}
}
