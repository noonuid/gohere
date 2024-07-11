package problem0061

import "github.com/noonuid/go/leetcode/structure"

// 61. 旋转链表

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	tail := head
	n := 1
	for tail.Next != nil {
		n++
		tail = tail.Next
	}
	k = k % n
	if k == 0 {
		return head
	}
	reTail := head
	for i := 1; i <= n-k-1; i++ {
		reTail = reTail.Next
	}
	reHead := reTail.Next
	reTail.Next = nil
	tail.Next = head
	return reHead
}
