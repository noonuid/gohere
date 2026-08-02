package problem0763

func partitionLabels(s string) []int {
	result := []int{}
	lastPos := [26]int{}
	for i, ch := range s {
		lastPos[ch-'a'] = i
	}
	start, end := 0, 0
	for i, ch := range s {
		if end < lastPos[ch-'a'] {
			end = lastPos[ch-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = end + 1
		}
	}
	return result
}
