package problem0287

// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。
// 假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。
// 你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
// 使用快慢指针求解。
func findDuplicate(nums []int) int {
	// 将 0 看作头节点，以此为起点，slow 每次移动一步，fast 每次移动两步。
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 记 a 为环外长度，b 为环的长度。
	// 第一次相遇时 slow 移动的长度是 nb，再使 slow 移动 a 个位置，slow 会移动到环的入口处。
	// 使 fast 从头节点移动 a 个位置到环的入口处，slow 和 fast 将会相遇。
	fast = 0
	for slow != fast {
		// slow 每次移动一步，fast 每次移动一步。
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// 使用二分查找求解。
func findDuplicateBinarySearch(nums []int) int {
	low, high := 1, len(nums)-1
	for low < high {
		mid := (low + high) >> 1
		count := 0
		for _, value := range nums {
			if value <= mid {
				count++
			}
		}
		// 对于 mid ∈ [1,target−1]，count <= mid，
		// 对于 mid ∈ [target,n]，count > mid。
		if count <= mid {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

// 暴力枚举。
// 超出时间限制。
func findDuplicateEnum(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return 0
}

// 使用 map 存储已遍历元素。
// 空间复杂度 O(n)，不符合题目要求。
func findDuplicateMap(nums []int) int {
	m := make(map[int]struct{}, len(nums)-1)
	for _, value := range nums {
		if _, ok := m[value]; ok {
			return value
		} else {
			m[value] = struct{}{}
		}
	}
	return 0
}
