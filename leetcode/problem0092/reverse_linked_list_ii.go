package problem0092

import "github.com/noonuid/go/leetcode/structure"

// 92. 反转链表 II

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	leftPre := dummyHead
	for i := 1; i < left; i++ {
		leftPre = leftPre.Next
	}

	cur, next := leftPre.Next, leftPre.Next.Next
	for i := 0; i < right-left; i++ {
		cur.Next = next.Next
		next.Next = leftPre.Next
		leftPre.Next = next
		next = cur.Next
	}

	return dummyHead.Next
}

func reverseBetween_concatenate(head *ListNode, left int, right int) *ListNode {
	var firstPre *ListNode
	first := head
	i := 1
	for ; i < left; i++ {
		firstPre = first
		first = first.Next
	}

	var pre *ListNode
	cur, next := first, first.Next
	for ; i <= right; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	first.Next = next
	if firstPre != nil {
		firstPre.Next = pre
		return head
	} else {
		return pre
	}
}
