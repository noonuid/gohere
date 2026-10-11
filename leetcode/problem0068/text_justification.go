package problem0068

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}
	n := len(words)
	for left := 0; left < n; {
		width := len(words[left])
		right := left + 1
		for right < n && width+1+len(words[right]) <= maxWidth {
			width = width + 1 + len(words[right])
			right++
		}
		var ln string
		if right-left == 1 || right == n {
			ln = strings.Join(words[left:right], " ") + strings.Repeat(" ", maxWidth-width)
		} else {
			numGaps := right - left - 1
			numSpaces := maxWidth - width + numGaps
			average, remainder := numSpaces/(numGaps), numSpaces%(numGaps)
			ln = strings.Join(words[left:left+remainder+1], strings.Repeat(" ", average+1))
			ln += strings.Repeat(" ", average)
			ln += strings.Join(words[left+remainder+1:right], strings.Repeat(" ", average))
		}
		result = append(result, ln)
		left = right
	}
	return result
}
