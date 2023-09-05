package problem0148

import (
	"sort"

	"github.com/noonuid/go/leetcode/structure"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

// 自底向上归并排序。
func sortList_bottom_up_merge_sort(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	merge := func(headOne, headTwo *ListNode) *ListNode {
		head := &ListNode{}
		node, nodeOne, nodeTwo := head, headOne, headTwo
		for nodeOne != nil && nodeTwo != nil {
			if nodeOne.Val < nodeTwo.Val {
				node.Next = nodeOne
				nodeOne = nodeOne.Next
			} else {
				node.Next = nodeTwo
				nodeTwo = nodeTwo.Next
			}
			node = node.Next
		}
		if nodeOne != nil {
			node.Next = nodeOne
		}
		if nodeTwo != nil {
			node.Next = nodeTwo
		}
		return head.Next
	}

	// 获取链表长度。
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}

	// 形式的头节点，位于真实的头节点前面。
	dummyHead := &ListNode{
		Next: head,
	}
	for subLen := 1; subLen < n; subLen <<= 1 {
		pre, cur := dummyHead, dummyHead.Next
		for cur != nil {
			headOne := cur
			for i := 1; i < subLen && cur.Next != nil; i++ {
				cur = cur.Next
			}

			headTwo := cur.Next
			cur.Next = nil
			cur = headTwo
			if cur != nil {
				for i := 1; i < subLen && cur.Next != nil; i++ {
					cur = cur.Next
				}

				next := cur.Next
				cur.Next = nil
				cur = next
			}

			pre.Next = merge(headOne, headTwo)
			for pre.Next != nil {
				pre = pre.Next
			}
		}
	}
	return dummyHead.Next
}

// 自顶向下归并排序。
func sortList_top_down_merge_sort(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	merge := func(headOne, headTwo *ListNode) *ListNode {
		head := &ListNode{}
		node, nodeOne, nodeTwo := head, headOne, headTwo
		for nodeOne != nil && nodeTwo != nil {
			if nodeOne.Val < nodeTwo.Val {
				node.Next = nodeOne
				nodeOne = nodeOne.Next
			} else {
				node.Next = nodeTwo
				nodeTwo = nodeTwo.Next
			}
			node = node.Next
		}
		if nodeOne != nil {
			node.Next = nodeOne
		}
		if nodeTwo != nil {
			node.Next = nodeTwo
		}
		return head.Next
	}

	var sortFn func(head, tail *ListNode) *ListNode
	sortFn = func(head, tail *ListNode) *ListNode {
		if head.Next == tail {
			head.Next = nil
			return head
		}

		slow, quick := head, head
		for quick != tail && quick.Next != tail {
			slow, quick = slow.Next, quick.Next.Next
		}
		mid := slow
		headOne := sortFn(head, mid)
		headTwo := sortFn(mid, tail)
		res := merge(headOne, headTwo)
		return res
	}

	return sortFn(head, nil)
}

// 转换为数组，然后排序。
func sortList_convert_to_array(head *ListNode) *ListNode {
	nums := []int{}
	for node := head; node != nil; node = node.Next {
		nums = append(nums, node.Val)
	}
	numsLen := len(nums)
	if numsLen > 0 {
		sort.Ints(nums)
		head = &ListNode{
			Val: nums[0],
		}
	} else {
		head = nil
	}
	tail := head
	for i := 1; i < numsLen; i++ {
		tail.Next = &ListNode{
			Val: nums[i],
		}
		tail = tail.Next
	}
	return head
}

// 冒泡排序。
// 超出时间限制。
func sortList_bubble(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	for outer := head.Next; outer != nil; outer = outer.Next {
		cur, next := head, head.Next
		stop := true
		for next != nil {
			if cur.Val > next.Val {
				cur.Val, next.Val = next.Val, cur.Val
				stop = false
			}
			cur, next = next, next.Next
		}
		if stop {
			break
		}
	}

	return head
}
