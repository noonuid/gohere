package nc156

import "sort"

// 排序解法
func foundOnceNumber(arr []int, k int) int {
	// write code here
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			i += k - 1
		} else {
			return arr[i]
		}
	}

	return arr[len(arr)-1]
}

// 使用 Hash Map
func foundOnceNumberHashMap(arr []int, k int) int {
	// write code here
	var store map[int]bool = make(map[int]bool)
	for _, num := range arr {
		if _, ok := store[num]; ok {
			store[num] = true
		} else {
			store[num] = false
		}
	}

	for key, value := range store {
		if !value {
			return key
		}
	}

	return 0
}
