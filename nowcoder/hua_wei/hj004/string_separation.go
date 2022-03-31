package hj004

import "fmt"

func Do() {
	var input string
	var number, _ = fmt.Scan(&input)
	for number > 0 {
		for len(input) >= 8 {
			fmt.Println(input[:8])
			input = input[8:]
		}
		if len(input) > 0 {
			for len(input) < 8 {
				input += "0"
			}
			fmt.Println(input)
		}

		number, _ = fmt.Scan(&input)
	}
}
