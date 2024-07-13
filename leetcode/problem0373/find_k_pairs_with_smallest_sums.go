package problem0373

import (
	"container/heap"
	"sort"
)

// 373. 查找和最小的 K 对数字

// 给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

// 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

// 堆。
func kSmallestPairs_heap(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	answer := make([][]int, 0, k)
	hp := &minHeap{
		data:  make([]pair, 0, k),
		nums1: nums1,
		nums2: nums2,
	}
	for i := 0; i < k && i < n1; i++ {
		hp.data = append(hp.data, pair{i: i, j: 0})
	}
	for hp.Len() > 0 && len(answer) < k {
		top := heap.Pop(hp).(pair)
		answer = append(answer, []int{nums1[top.i], nums2[top.j]})
		if top.j+1 < k && top.j+1 < n2 {
			heap.Push(hp, pair{top.i, top.j + 1})
		}
	}
	return answer
}

type pair struct{ i, j int }

type minHeap struct {
	nums1 []int
	nums2 []int
	data  []pair
}

func (h *minHeap) Push(x any) {
	h.data = append(h.data, x.(pair))
}

func (h *minHeap) Pop() any {
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return last
}

func (h *minHeap) Len() int {
	return len(h.data)
}

func (h *minHeap) Less(i int, j int) bool {
	return h.nums1[h.data[i].i]+h.nums2[h.data[i].j] < h.nums1[h.data[j].i]+h.nums2[h.data[j].j]
}

func (h *minHeap) Swap(i int, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// 暴力枚举。
func kSmallestPairs_enum(nums1 []int, nums2 []int, k int) [][]int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n1, n2 := min(len(nums1), k), min(len(nums2), k)
	answer := make([][]int, 0, n1*n2)
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			answer = append(answer, []int{nums1[i], nums2[j]})
		}
	}
	sort.Slice(answer, func(i, j int) bool {
		return (answer[i][0] + answer[i][1]) < (answer[j][0] + answer[j][1])
	})
	return answer[:k]
}
