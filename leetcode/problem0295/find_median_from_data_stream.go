package problem0295

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	pqLt priorityQueue
	pqGt priorityQueue
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	pqLt, pqGt := &this.pqLt, &this.pqGt
	if pqLt.Len() == 0 || num <= -pqLt.IntSlice[0] {
		heap.Push(pqLt, -num)
		if pqLt.Len()-1 > pqGt.Len() {
			heap.Push(pqGt, -heap.Pop(pqLt).(int))
		}
	} else {
		heap.Push(pqGt, num)
		if pqGt.Len() > pqLt.Len() {
			heap.Push(pqLt, -heap.Pop(pqGt).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.pqLt.Len() > this.pqGt.Len() {
		return float64(-this.pqLt.IntSlice[0])
	}
	return float64(this.pqGt.IntSlice[0]-this.pqLt.IntSlice[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type priorityQueue struct {
	sort.IntSlice
}

func (pq *priorityQueue) Push(x any) {
	pq.IntSlice = append(pq.IntSlice, x.(int))
}

func (pq *priorityQueue) Pop() any {
	n := len(pq.IntSlice)
	if n == 0 {
		return nil
	}
	x := pq.IntSlice[n-1]
	pq.IntSlice = pq.IntSlice[:n-1]
	return x
}
