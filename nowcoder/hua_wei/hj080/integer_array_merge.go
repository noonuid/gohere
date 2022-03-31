package hj080

import (
	"fmt"
	"sort"
)

func Do() {
	var n1, n2, scanN int
	var nums1, nums2 []int

	fmt.Scan(&n1)
	nums1 = make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&nums1[i])
	}
	fmt.Scan(&n2)
	nums2 = make([]int, n2)
	for i := 0; i < n2; i++ {
		scanN, _ = fmt.Scan(&nums2[i])
	}

	for scanN > 0 {
		sort.Ints(nums1)
		sort.Ints(nums2)
		var result = make([]int, 0)
		var i, j = 0, 0
		for i < n1 && j < n2 {
			if nums1[i] < nums2[j] {
				result = append(result, nums1[i])
				i++
			} else {
				result = append(result, nums2[j])
				j++
			}
		}
		for i < n1 {
			result = append(result, nums1[i])
			i++
		}
		for j < n2 {
			result = append(result, nums2[j])
			j++
		}

		fmt.Print(result[0])
		for i := 1; i < len(result); i++ {
			if result[i] > result[i-1] {
				fmt.Print(result[i])
			}
		}
		fmt.Print("\n")

		fmt.Scan(&n1)
		nums1 = make([]int, n1)
		for i := 0; i < n1; i++ {
			fmt.Scan(&nums1[i])
		}
		fmt.Scan(&n2)
		nums2 = make([]int, n2)
		for i := 0; i < n2; i++ {
			scanN, _ = fmt.Scan(&nums2[i])
		}
	}
}
