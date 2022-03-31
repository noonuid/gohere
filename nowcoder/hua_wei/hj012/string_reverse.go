package hj012

import (
	"bufio"
	"fmt"
	"os"
)

func Do() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()

	head, tail := 0, len(line)-1
	for head < tail {
		line[head], line[tail] = line[tail], line[head]
		head++
		tail--
	}
	fmt.Println(string(line))
}
