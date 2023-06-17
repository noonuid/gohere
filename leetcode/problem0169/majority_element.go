package problem0169

// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// Boyer-Moore 投票算法。
func majorityElementBoyerMoore(nums []int) int {
	candidate, count := -1, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// 使用 map 存储频率。
func majorityElementMap(nums []int) int {
	m := make(map[int]int, len(nums)/2+1)
	for _, num := range nums {
		m[num]++
		if m[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}

// 暴力查找。
func majorityElementBruteForce(nums []int) int {
	for i := 0; i < len(nums)/2+1; i++ {
		cur := nums[i]
		count := 1
		for j := i + 1; j < len(nums); j++ {
			if cur == nums[j] {
				count++
			}
		}
		if count > len(nums)/2 {
			return cur
		}
	}
	return -1
}
