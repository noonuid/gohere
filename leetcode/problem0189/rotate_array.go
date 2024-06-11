package problem0189

// 189. 轮转数组

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

// 数组翻转.
func rotate_reverse(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}

	reverse := func(nums []int) {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

// 环状替换。
func rotate_circular_replacement(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}

	for start, count := 0, 0; count < n; start++ {
		temp := nums[start]
		cur, right := -1, (start+k)%n
		for cur != start {
			nums[right], temp = temp, nums[right]
			cur, right = right, (right+k)%n
			count++
		}
	}
}

// 先将元素按指定顺序复制到另一个数组，再复制回来。
func rotate_copy(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}

	result := make([]int, 0, len(nums))
	result = append(result, nums[len(nums)-k:]...)
	result = append(result, nums[:len(nums)-k]...)
	copy(nums, result)
}
