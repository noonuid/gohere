package problem0025

import "github.com/noonuid/go/leetcode/structure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	preSubTail := dummy
	subHead := head
	for subHead != nil {
		subTail := preSubTail
		for range k {
			subTail = subTail.Next
			if subTail == nil {
				return dummy.Next
			}
		}
		pre, cur, next := subTail.Next, subHead, subHead.Next
		for pre != subTail {
			next = cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		preSubTail.Next = subTail
		preSubTail = subHead
		subHead = cur
	}
	return dummy.Next
}
