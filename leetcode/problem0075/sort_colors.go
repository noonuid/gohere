package problem0075

// 75. 颜色分类

// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

// 必须在不使用库内置的 sort 函数的情况下解决这个问题。

// 双指针。
func sortColors_double_pointer(nums []int) {
	n := len(nums)
	p0, p1 := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0, p1 = p0+1, p1+1
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

// 单指针。
func sortColors_single_pointer(nums []int) {
	n := len(nums)
	cur := 0
	// 将所有的 0 移动到前面。
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[cur], nums[i] = nums[i], nums[cur]
			cur++
		}
	}
	// 将所有的 1 移动到连续的 0 的后面。
	for i := cur; i < n; i++ {
		if nums[i] == 1 {
			nums[cur], nums[i] = nums[i], nums[cur]
			cur++
		}
	}
	// 两轮移动之后，所有的 2 在最后面。
}

// 冒泡排序。
func sortColors_bubble(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
