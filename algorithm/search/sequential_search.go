package search

func SequentialSearch(arr []int, target int) int {
	for i, item := range arr {
		if item == target {
			return i
		}
	}

	return -1
}
