package hj005

import (
	"fmt"
	"strconv"
)

func Do() {
	var input string
	var scanN, _ = fmt.Scan(&input)
	for scanN > 0 {
		var output, _ = strconv.ParseInt(input, 0, 0)
		fmt.Println(output)

		scanN, _ = fmt.Scan(&input)
	}
}
