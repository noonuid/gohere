package problem0018

import (
	"sort"
)

// 18. 四数之和

// 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的
// 四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target

// 你可以按 任意顺序 返回答案 。

// 双指针。
func fourSum_double_pointer(nums []int, target int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n < 4 {
		return res
	}
	sort.Ints(nums)
	for a := 0; a < n-3 && nums[a]+nums[a+1]+nums[a+2]+nums[a+3] <= target; a++ {
		if a > 0 && nums[a] == nums[a-1] || nums[a]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for b := a + 1; b < n-2 && nums[a]+nums[b]+nums[b+1]+nums[b+2] <= target; b++ {
			if b > a+1 && nums[b] == nums[b-1] || nums[a]+nums[b]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := b+1, n-1; left < right; {
				if sum := nums[a] + nums[b] + nums[left] + nums[right]; sum == target {
					res = append(res, []int{nums[a], nums[b], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
				} else {
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				}
			}
		}
	}
	return res
}

// 暴力枚举。
func fourSum_brute(nums []int, target int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n < 4 {
		return res
	}
	sort.Ints(nums)
	for a := 0; a < n-3 && nums[a]+nums[a+1]+nums[a+2]+nums[a+3] <= target; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		for b := a + 1; b < n-2 && nums[a]+nums[b]+nums[b+1]+nums[b+2] <= target; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for c := b + 1; c < n-1 && nums[a]+nums[b]+nums[c]+nums[c+1] <= target; c++ {
				if c > b+1 && nums[c] == nums[c-1] {
					continue
				}
				for d := c + 1; d < n && nums[a]+nums[b]+nums[c]+nums[d] <= target; d++ {
					if d > c+1 && nums[d] == nums[d-1] {
						continue
					}
					if nums[a]+nums[b]+nums[c]+nums[d] == target {
						res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					}
				}
			}
		}
	}
	return res
}
