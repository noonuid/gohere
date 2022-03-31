package i

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Do() {
	var reader = bufio.NewReader(os.Stdin)

	var line, _, _ = reader.ReadLine()
	for len(line) > 0 {
		var strs = strings.Split(string(line), " ")
		sort.Strings(strs)
		fmt.Println(strings.Join(strs, " "))

		line, _, _ = reader.ReadLine()
	}
}
