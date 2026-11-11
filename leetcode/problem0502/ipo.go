package problem0502

import (
	"container/heap"
	"sort"
)

type profitHeap struct {
	sort.IntSlice
}

func (h profitHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *profitHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *profitHeap) Pop() any {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	type pair struct{ capital, profit int }
	pairs := make([]pair, n)
	for i, p := range profits {
		pairs[i] = pair{capital[i], p}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].capital < pairs[j].capital })

	h := &profitHeap{}
	cur := 0
	for i := 0; i < k; i++ {
		for ; cur < n && w >= pairs[cur].capital; cur++ {
			heap.Push(h, pairs[cur].profit)
		}
		if h.Len() <= 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}
