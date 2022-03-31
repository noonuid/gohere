package hj013

import "fmt"

func Do() {
	var word string
	var words []string = make([]string, 0)

	var scanN, _ = fmt.Scan(&word)
	for scanN > 0 {
		words = append(words, word)

		scanN, _ = fmt.Scan(&word)
	}

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Printf("%s ", words[i])
	}
}
