package problem0234

import "github.com/noonuid/go/nowcoder/structure"

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structure.ListNode

// 先反转后半段链表再对比前后两段链表的值。
func isPalindrome_reverse(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转后半段链表。
	var last, next *ListNode
	for cur := slow.Next; cur != nil; {
		next = cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	// 比较两段链表是否相等。
	for first, second := head, last; first != nil && second != nil; first, second = first.Next, second.Next {
		if first.Val != second.Val {
			return false
		}
	}
	return true
}

// 将值复制到数组中再用双指针法。
func isPalindrome_copy(head *ListNode) bool {
	values := []int{}
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		if values[i] != values[j] {
			return false
		}
	}
	return true
}
