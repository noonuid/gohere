package problem0274

import "sort"

// 274. H 指数

// 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

// 根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。

// 对引用计数。
func hIndex_counter(citations []int) int {
	n := len(citations)
	counter := make([]int, n+1)
	for i := 0; i < n; i++ {
		if citations[i] >= n {
			counter[n]++
		} else {
			counter[citations[i]]++
		}
	}
	for i, total := n, 0; i > -1; i-- {
		total = total + counter[i]
		if total >= i {
			return i
		}
	}
	return 0
}

// 排序。
func hIndex_sort(citations []int) int {
	sort.Ints(citations)
	h := 0
	for i := len(citations) - 1; i > -1 && citations[i] > h; i-- {
		h++
	}
	return h
}
