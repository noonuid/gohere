package problem0002

import "github.com/noonuid/go/leetcode/structure"

// 2. 两数相加

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	node1, node2 := l1, l2
	sum, carry := 0, 0
	for node1 != nil || node2 != nil || carry > 0 {
		sum = carry
		if node1 != nil {
			sum += node1.Val
			node1 = node1.Next
		}
		if node2 != nil {
			sum += node2.Val
			node2 = node2.Next
		}

		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	return head
}
