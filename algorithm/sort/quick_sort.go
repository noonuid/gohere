package sort

// 快速排序
func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	index := partition(nums)
	QuickSort(nums[:index])
	QuickSort(nums[index+1:])

	return nums
}

// 根据枢轴量划分切片，返回枢轴量在切片中的索引
// 该划分算法可以保证返回的索引处存储的一定是枢轴量
func partition(nums []int) int {
	// 选取枢轴量
	pivot := nums[len(nums)-1]

	// i 会始终指向在 j 之前的并且小于或等于枢轴量的最后一个元素
	i := -1
	for j := 0; j < len(nums); j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return i
}

// 快速排序，分区算法使用双指针
func QuickSortTwoPointer(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	index := partitionTwoPointer(nums)
	QuickSortTwoPointer(nums[:index])
	QuickSortTwoPointer(nums[index+1:])

	return nums
}

// 使用双指针根据枢轴量划分切片，返回枢轴量在切片中的索引
// 该划分算法可以保证返回的索引处存储的一定是枢轴量
func partitionTwoPointer(nums []int) int {
	low, high := 0, len(nums)-1
	pivot := nums[high]
	for low < high {
		for low < high && nums[low] <= pivot {
			low++
		}
		nums[high] = nums[low]
		for high > low && nums[high] >= pivot {
			high--
		}
		nums[low] = nums[high]
	}
	nums[low] = pivot
	return low
}
