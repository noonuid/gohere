package problem0215

// 方法一：先排序再返回
func findKthLargest(nums []int, k int) int {
	quick(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func quick(nums []int, low, high int) {
	if len(nums) <= 1 {
		return
	}

	if low < high {
		pivotIndex := partion(nums, low, high)
		quick(nums, low, pivotIndex-1)
		quick(nums, pivotIndex, high)
	}
}

func partion(nums []int, low, high int) int {
	var pivot = nums[(low+high)/2]

	for low <= high {
		for nums[low] < pivot {
			low++
		}
		for nums[high] > pivot {
			high--
		}

		if low <= high {
			nums[low], nums[high] = nums[high], nums[low]

			low++
			high--
		}
	}

	return low
}

// 方法二：在快速排序过程中，当枢轴量就是第 K 大的数时，直接返回，不必完成整个排序过程
func findKthLargestSelect(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, index int) int {
	var pivotIndex = partionSelect(nums, left, right)

	if pivotIndex == index {
		return nums[pivotIndex]
	} else if pivotIndex < index {
		return quickSelect(nums, pivotIndex+1, right, index)
	}
	return quickSelect(nums, left, pivotIndex-1, index)
}

func partionSelect(nums []int, left, right int) int {
	var pivot = nums[right]

	var i = left - 1
	for j := left; j <= right; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return i
}
