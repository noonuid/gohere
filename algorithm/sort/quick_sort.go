package sort

func QuickSort(nums []int) []int {
	quick(nums, 0, len(nums)-1)
	return nums
}

func quick(nums []int, low, high int) {
	if low < high {
		var index = partition(nums, low, high)

		quick(nums, low, index-1)
		quick(nums, index+1, high)
	}
}

// 根据枢轴量划分切片，返回枢轴量在切片中的索引
// 该划分算法可以保证返回的索引处存储的一定是枢轴量
func partition(nums []int, low, high int) int {
	// 选取枢轴量
	var pivot = nums[high]

	// i 会始终指向在 j 之前的并且小于枢轴量的最后一个元素
	var i = low - 1
	for j := low; j <= high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return i
}

// 第二种快速排序，与第一种的分区算法不同
func QuickSortI(nums []int) []int {
	quickI(nums, 0, len(nums)-1)

	return nums
}

func quickI(nums []int, low, high int) {
	if low < high {
		var index = partitionI(nums, low, high)

		quickI(nums, low, index-1)
		quickI(nums, index, high)
	}
}

// 该划分算法不能保证返回的索引处存储的一定是枢轴量
// 只能保证索引及以上为大于等于枢轴量的元素，索引以下为小于等于枢轴量的元素
func partitionI(nums []int, low, high int) int {
	// 选取枢轴量
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
