package hj007

import "fmt"

func Do() {
	var input float32
	var number, _ = fmt.Scan(&input)
	for number > 0 {
		var fraction = input - float32(int(input))
		var output = 0
		if fraction >= 0.5 {
			output = int(input) + 1
		} else {
			output = int(input)
		}
		fmt.Println(output)

		number, _ = fmt.Scan(&input)
	}
}
