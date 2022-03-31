package problem1859

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	result := make([]string, len(words))

	index := 0
	for i := 0; i < len(words); i++ {
		index, _ = strconv.Atoi(words[i][len(words[i])-1 : len(words[i])])
		result[index-1] = words[i][:len(words[i])-1]
	}

	return strings.Join(result, " ")
}
