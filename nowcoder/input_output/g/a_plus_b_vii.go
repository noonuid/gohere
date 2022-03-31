package g

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Do() {
	var reader = bufio.NewReader(os.Stdin)

	var line, _, _ = reader.ReadLine()
	for len(line) > 0 {
		var strs = strings.Split(string(line), " ")
		var sum = 0
		for _, str := range strs {
			var num, _ = strconv.Atoi(str)
			sum += num
		}
		fmt.Println(sum)

		line, _, _ = reader.ReadLine()
	}
}
