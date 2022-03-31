package hj002

import (
	"bufio"
	"fmt"
	"os"
)

// 方法一：字符串一次性读入
func Do() {
	var reader = bufio.NewReader(os.Stdin)
	var input string
	var character rune

	var line, _, _ = reader.ReadLine()
	input = string(line)
	line, _, _ = reader.ReadLine()
	for len(line) > 0 {
		character = rune(line[0])

		var count = 0
		character = toLower(character)
		for _, ch := range input {
			ch = toLower(ch)
			if ch == character {
				count++
			}
		}
		fmt.Println(count)

		line, _, _ = reader.ReadLine()
		input = string(line)
		line, _, _ = reader.ReadLine()
	}
}

// 将小写字母转换为大写字母，若传入字符不是小写字母，则返回原字母
func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	} else {
		return ch
	}
}

// 方法二：每次读入一个单词，直到读取完所有输入
func DoReadSingleWord() {
	var input []rune
	var character rune
	var word string

	var number, _ = fmt.Scan(&word)
	for number > 0 {
		input = append(input, []rune(word)...)

		number, _ = fmt.Scan(&word)
	}

	character = toLower(input[len(input)-1])
	input = input[:len(input)-1]

	var count = 0
	for _, ch := range input {
		if toLower(ch) == character {
			count++
		}
	}

	fmt.Println(count)
}
