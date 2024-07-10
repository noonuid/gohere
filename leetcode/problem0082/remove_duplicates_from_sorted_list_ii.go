package problem0082

import (
	"github.com/noonuid/go/leetcode/structure"
)

// 82. 删除排序链表中的重复元素 II

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headPre := &ListNode{Next: head}
	cur := headPre
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val != cur.Next.Next.Val {
			cur = cur.Next
		} else {
			node := cur.Next
			for cur.Next != nil && cur.Next.Val == node.Val {
				cur.Next = cur.Next.Next
			}
		}
	}
	return headPre.Next
}
