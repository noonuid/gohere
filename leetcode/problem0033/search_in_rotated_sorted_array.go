package problem0033

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			// mid 在左边升序序列中

			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// mid 在右边升序序列中

			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// 查找 target，先使用二分查找得到最小值的位置
func searchMin(nums []int, target int) int {
	// 寻找最大值和最小值的位置
	max, min := 0, len(nums)-1
	if nums[max] > nums[min] {
		for (min - max) > 1 {
			mid := (max + min) >> 1
			if nums[mid] >= nums[max] {
				max = mid
			} else {
				min = mid
			}
		}
	} else {
		min = max
	}

	// 将 nums 转化为单调升序数组
	sorted := append(nums[min:], nums[:min]...)
	low, high := 0, len(sorted)-1
	for low <= high {
		mid := (low + high) >> 1
		if sorted[mid] == target {
			if mid < len(nums)-min {
				return mid + min
			} else {
				return mid - (len(nums) - min)
			}
		} else if sorted[mid] > target {
			high = mid - 1
		} else if sorted[mid] < target {
			low = mid + 1
		}
	}

	return -1
}

// 查找 target，暴力枚举
func searchEnum(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}
