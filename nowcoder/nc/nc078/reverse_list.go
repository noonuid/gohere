package nc078

import (
	"github.com/noonuid/go/nowcoder/structure"
)

type ListNode = structure.ListNode

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var prev *ListNode
	for pHead != nil {
		next := pHead.Next
		pHead.Next = prev
		prev = pHead
		pHead = next
	}
	return prev
}
