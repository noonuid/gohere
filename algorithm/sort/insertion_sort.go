package sort

func InsertionSort(nums []int) []int {
	var length = len(nums)

	for i := 1; i < length; i++ {
		var cur = nums[i]

		j := i
		for j > 0 && nums[j-1] > cur {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = cur
	}

	return nums
}
