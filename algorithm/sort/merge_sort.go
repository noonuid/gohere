package sort

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	var middle = len(nums) / 2
	var left = MergeSort(nums[:middle])
	var right = MergeSort(nums[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result = make([]int, len(left)+len(right))
	var i, j = 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		result[i+j] = left[i]
		i++
	}
	for j < len(right) {
		result[i+j] = right[j]
		j++
	}

	return result
}
