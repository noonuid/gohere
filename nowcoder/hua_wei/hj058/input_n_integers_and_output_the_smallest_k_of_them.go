package hj058

import "fmt"

func Do() {
	var n, k = 0, 0
	var nScan, _ = fmt.Scan(&n, &k)
	for nScan > 0 {
		var nums = make([]int, n)
		// 读取整数数组
		for n > 0 {
			fmt.Scan(&nums[len(nums)-n])
			n--
		}

		quick(nums, 0, len(nums)-1, k-1)
		var result = nums[:k]
		for _, num := range result {
			fmt.Printf("%v ", num)
		}
		fmt.Print("\n")

		nScan, _ = fmt.Scan(&n, &k)
	}
}

// 快速排序方法
func quick(nums []int, low, high, index int) {
	if len(nums) <= 1 {
		return
	}

	if low < high {
		var pivotIndex = partion(nums, low, high)

		if pivotIndex == index {
			quick(nums[:pivotIndex+1], 0, pivotIndex, pivotIndex+1)
		} else {
			if low < high {
				quick(nums, low, pivotIndex-1, index)
				quick(nums, pivotIndex, high, index)
			}
		}
	}
}

func partion(nums []int, low, high int) int {
	var pivot = nums[high]

	var i = low - 1
	for j := low; j <= high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return i
}
