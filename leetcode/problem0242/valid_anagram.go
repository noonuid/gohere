package problem0242

import "sort"

func isAnagram(s string, t string) bool {
	bytes1 := []byte(s)
	bytes2 := []byte(t)

	sort.Slice(bytes1, func(i, j int) bool { return bytes1[i] < bytes1[j] })
	sort.Slice(bytes2, func(i, j int) bool { return bytes2[i] < bytes2[j] })

	return string(bytes1) == string(bytes2)
}
