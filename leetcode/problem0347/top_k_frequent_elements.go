package problem0347

import (
	"sort"
)

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

func topKFrequent(nums []int, k int) []int {
	// fs 记录各元素的频率。
	fs := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		fs[nums[i]]++
	}
	diffNums := make([]int, len(fs))
	index := 0
	for num := range fs {
		diffNums[index] = num
		index++
	}
	sort.Slice(diffNums, func(i, j int) bool { return fs[diffNums[i]] > fs[diffNums[j]] })
	return diffNums[:k]
}
