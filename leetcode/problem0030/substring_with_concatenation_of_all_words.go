package problem0030

import (
	"slices"
)

func findSubstring(s string, words []string) []int {
	result := []int{}
	slices.Sort(words)
	lenS, lenWords, lenWord := len(s), len(words), len(words[0])
	left, right := 0, lenWords*lenWord-1
	windowWords := make([]string, lenWords)
	for right < lenS {
		for i := left; i <= right; i = i + lenWord {
			windowWords[(i-left)/lenWord] = s[i : i+lenWord]
		}
		slices.Sort(windowWords)
		if slices.Equal(words, windowWords) {
			result = append(result, left)
		}
		left, right = left+1, right+1
	}
	return result
}

func findSubstring_map(s string, words []string) []int {
	result := []int{}
	lenS, lenWords, lenWord := len(s), len(words), len(words[0])
	for i := 0; i < lenWord && i <= lenS-lenWords*lenWord; i++ {
		differ := make(map[string]int)
		for j := 0; j < lenWords; j++ {
			differ[s[i+j*lenWord:i+(j+1)*lenWord]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		if len(differ) == 0 {
			result = append(result, i)
		}
		for left := i + lenWord; left <= lenS-lenWords*lenWord; left = left + lenWord {
			word := s[left-lenWord : left]
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
			word = s[left+(lenWords-1)*lenWord : left+lenWords*lenWord]
			differ[word]++
			if differ[word] == 0 {
				delete(differ, word)
			}
			if len(differ) == 0 {
				result = append(result, left)
			}
		}
	}
	return result
}
