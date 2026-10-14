package problem0057

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[1] < left {
			result = append(result, interval)
		} else if right < interval[0] {
			if !merged {
				result = append(result, []int{left, right})
				merged = true
			}
			result = append(result, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		result = append(result, []int{left, right})
	}
	return result
}
