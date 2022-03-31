package hj001

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 方法一：一次读入一整行
func Do() {
	var reader = bufio.NewReader(os.Stdin)
	var input, _, _ = reader.ReadLine()
	for len(input) > 0 {
		fmt.Println(lengthOfTheLastWord(string(input)))

		input, _, _ = reader.ReadLine()
	}
}

func lengthOfTheLastWord(str string) int {
	var strs = strings.Split(str, " ")

	return len(strs[len(strs)-1])
}

// 方法二：一次读入一个单词，直到读入最后一个单词
func DoReadSingleWord() {
	var word string
	var number, _ = fmt.Scan(&word)
	for number > 0 {
		number, _ = fmt.Scan(&word)
	}
	fmt.Println(len(word))
}
