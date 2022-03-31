package nc088

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, K int) int {
	// write code here
	quick(a, 0, len(a)-1)
	return a[len(a)-K]
}

func quick(nums []int, low, high int) {
	if len(nums) <= 1 {
		return
	}

	if low < high {
		var pivotIndex = partion(nums, low, high)

		quick(nums, low, pivotIndex-1)
		quick(nums, pivotIndex, high)
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
