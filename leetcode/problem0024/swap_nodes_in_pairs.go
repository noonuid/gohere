package problem0024

import "github.com/noonuid/go/leetcode/structure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func swapPairs_iteration(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	preNode2 := dummy
	for preNode2.Next != nil && preNode2.Next.Next != nil {
		node1, node2 := preNode2.Next, preNode2.Next.Next
		node1.Next, node2.Next = node2.Next, node1
		preNode2.Next = node2
		preNode2 = node1
	}
	return dummy.Next
}
