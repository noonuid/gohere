package problem0023

import "github.com/noonuid/go/leetcode/structure"

// 23. 合并 K 个升序链表

// 给你一个链表数组，每个链表都已经按升序排列。

// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

// 分治合并。
func mergeKLists_divide_and_conquer(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}

	var merge func(left, right int) *ListNode
	merge = func(left, right int) *ListNode {
		if left == right {
			return lists[left]
		}

		mid := (left + right) >> 1
		first := merge(left, mid)
		second := merge(mid+1, right)

		// 合并 first 与 second。
		if first == nil {
			return second
		} else if second == nil {
			return first
		}
		head := &ListNode{}
		tail := head
		i, j := first, second
		for i != nil && j != nil {
			if i.Val < j.Val {
				tail.Next = i
				i, tail = i.Next, tail.Next
			} else {
				tail.Next = j
				j, tail = j.Next, tail.Next
			}
		}
		if i != nil {
			tail.Next = i
		} else {
			tail.Next = j
		}
		return head.Next
	}

	return merge(0, k-1)
}

// 依次合并。
func mergeKLists_sequentially(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	mergeTwoLists := func(first, second *ListNode) *ListNode {
		if first == nil {
			return second
		} else if second == nil {
			return first
		}
		head := &ListNode{}
		tail := head
		i, j := first, second
		for i != nil && j != nil {
			if i.Val < j.Val {
				tail.Next = i
				i, tail = i.Next, tail.Next
			} else {
				tail.Next = j
				j, tail = j.Next, tail.Next
			}
		}
		if i != nil {
			tail.Next = i
		} else {
			tail.Next = j
		}
		return head.Next
	}
	first := lists[0]
	for i := 1; i < k; i++ {
		second := lists[i]
		first = mergeTwoLists(first, second)
	}
	return first
}
