package j

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var input string
	var reader = bufio.NewReader(os.Stdin)

	var bytes, _, _ = reader.ReadLine()
	for len(bytes) > 0 {
		input = string(bytes)
		strs := strings.Split(input, ",")
		sort.Strings(strs)
		fmt.Println(strings.Join(strs, ","))

		bytes, _, _ = reader.ReadLine()
	}
}
