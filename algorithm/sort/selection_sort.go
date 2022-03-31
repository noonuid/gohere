package sort

func SelectionSort(nums []int) []int {
	var length = len(nums)

	for i := 0; i < length-1; i++ {
		var min = i
		for j := i + 1; j < length; j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
		}
	}

	return nums
}
