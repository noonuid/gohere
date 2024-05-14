package problem0167

// 双指针。
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if sum := numbers[left] + numbers[right]; sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{0, 0}
}

// 二分查找。
func twoSum_binary(numbers []int, target int) []int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		low, high := i+1, n-1
		t := target - numbers[i]
		for low <= high {
			mid := low + (high-low)/2
			if numbers[mid] < t {
				low = mid + 1
			} else if numbers[mid] > t {
				high = mid - 1
			} else {
				return []int{i + 1, mid + 1}
			}
		}
	}
	return []int{0, 0}
}
