package problem0128

import "sort"

// 128. 最长连续序列

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// 使用哈希表记录数组中元素。
func longestConsecutive_map(nums []int) int {
	max := 0
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	for _, num := range nums {
		// num 的前驱 num-1 不存在时，num 可以作为一段连续序列的起点。
		if !m[num-1] {
			curLen, curNum := 1, num
			for m[curNum+1] {
				curLen++
				curNum++
			}
			if curLen > max {
				max = curLen
			}
		}
	}
	return max
}

// 先将数组排序，再寻找序列。
func longestConsecutive_sort(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	sort.Ints(nums)
	max := 1
	for begin := 0; begin < n-max; begin++ {
		repeat := 0
		for end := begin + 1; end < n; end++ {
			if diff := nums[end] - nums[end-1]; diff != 1 {
				if curLen := end - begin - repeat; curLen > max {
					max = curLen
				}
				if diff == 0 {
					repeat++
					continue
				} else {
					break
				}
			} else if end == n-1 {
				// 数组的最后一位元素属于连续序列。
				if curLen := end - begin + 1 - repeat; curLen > max {
					max = curLen
				}
			}
		}
	}
	return max
}
