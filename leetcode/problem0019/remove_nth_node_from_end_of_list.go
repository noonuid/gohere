package problem0019

import (
	"github.com/noonuid/go/leetcode/structure"
)

// 19. 删除链表的倒数第 N 个结点

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

// 使用双指针一次遍历。
func removeNthFromEnd_double_pointer(head *ListNode, n int) *ListNode {
	reHead := &ListNode{Val: 0, Next: head}
	first, second := head, reHead
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first != nil {
		first, second = first.Next, second.Next
	}
	second.Next = second.Next.Next
	return reHead.Next
}

// 使用栈存储遍历结果。
func removeNthFromEnd_stack(head *ListNode, n int) *ListNode {
	reHead := &ListNode{Val: 0, Next: head}
	stack := []*ListNode{}
	for node := reHead; node != nil; node = node.Next {
		stack = append(stack, node)
	}
	nthPre := stack[len(stack)-1-n]
	nthPre.Next = nthPre.Next.Next
	return reHead.Next
}

// 两次遍历。
func removeNthFromEnd_two_traversal(head *ListNode, n int) *ListNode {
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	reHead := &ListNode{
		Next: head,
	}
	nthPre := reHead
	for i := 0; i < length-n; i++ {
		nthPre = nthPre.Next
	}
	nthPre.Next = nthPre.Next.Next
	return reHead.Next
}
