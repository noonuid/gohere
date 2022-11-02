package problem0021

import (
	"github.com/noonuid/go/leetcode/structure"
)

// Definition for singly-linked list.
type ListNode = structure.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode

	var node *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}

		if tail == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}

	if l1 != nil {
		if tail == nil {
			head = l1
		} else {
			tail.Next = l1
		}
	} else if l2 != nil {
		if tail == nil {
			head = l2
		} else {
			tail.Next = l2
		}
	}

	return head
}
