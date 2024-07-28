package problem0086

import "github.com/noonuid/go/leetcode/structure"

// 86. 分隔链表

// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

// 你应当 保留 两个分区中每个节点的初始相对位置。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

// 将原始链表拆分为两个链表，再将两个链表拼接起来。
func partition(head *ListNode, x int) *ListNode {
	smallHead, largeHead := &ListNode{}, &ListNode{}
	small, large := smallHead, largeHead
	for node := head; node != nil; node = node.Next {
		if node.Val < x {
			small.Next = node
			small = small.Next
		} else {
			large.Next = node
			large = large.Next
		}
	}
	small.Next = largeHead.Next
	large.Next = nil
	return smallHead.Next
}

// 将小于 x 的节点插入到大于等于 x 的节点前面。
func partition_insertion(head *ListNode, x int) *ListNode {
	reHead := &ListNode{Val: x - 1, Next: head}
	left := reHead
	for ; left.Next != nil && left.Next.Val < x; left = left.Next {
	}
	pre, right := left, left.Next
	for right != nil {
		if right.Val < x {
			pre.Next = right.Next
			right.Next = left.Next
			left.Next = right

			left = left.Next
			right = pre.Next
		} else {
			pre = right
			right = right.Next
		}
	}
	return reHead.Next
}
