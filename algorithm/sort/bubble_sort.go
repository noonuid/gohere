package sort

// 冒泡排序
func BubbleSort(nums []int) []int {
	var length = len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
