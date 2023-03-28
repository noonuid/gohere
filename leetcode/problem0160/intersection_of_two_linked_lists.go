package problem0160

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for ptrA := headA; ptrA != nil; ptrA = ptrA.Next {
		for ptrB := headB; ptrB != nil; ptrB = ptrB.Next {
			if ptrA == ptrB {
				return ptrA
			}
		}
	}

	return nil
}

func getIntersectionNodeMap(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]interface{})
	for ptrA := headA; ptrA != nil; ptrA = ptrA.Next {
		m[ptrA] = nil
	}
	for ptrB := headB; ptrB != nil; ptrB = ptrB.Next {
		if _, ok := m[ptrB]; ok {
			return ptrB
		}
	}

	return nil
}

// 先使链表 A 和 B 剩余长度相等再找交叉点
func getIntersectionNodeAlign(headA, headB *ListNode) *ListNode {
	lengthA, lengthB := 0, 0
	ptrA, ptrB := headA, headB
	// 获取链表 A 和 B 的长度
	for ; ptrA != nil; ptrA = ptrA.Next {
		lengthA++
	}
	for ; ptrB != nil; ptrB = ptrB.Next {
		lengthB++
	}

	// 使链表 A 和 B 剩余长度相等
	ptrA, ptrB = headA, headB
	if lengthA >= lengthB {
		diff := lengthA - lengthB
		for ; diff > 0; diff-- {
			ptrA = ptrA.Next
		}
	} else {
		diff := lengthB - lengthA
		for ; diff > 0; diff-- {
			ptrB = ptrB.Next
		}
	}

	// 寻找交叉点
	for ptrA != nil && ptrB != nil {
		if ptrA == ptrB {
			return ptrA
		}
		ptrA, ptrB = ptrA.Next, ptrB.Next
	}

	return nil
}

// 先使链表 A 和 B 剩余长度相等再找交叉点
// 整个过程中，A 会连接到 B 后面，B 会连接到 A 后面
func getIntersectionNodeJoin(headA, headB *ListNode) *ListNode {
	ptrA, ptrB := headA, headB
	for ptrA != ptrB {
		if ptrA != nil {
			ptrA = ptrA.Next
		} else {
			ptrA = headB
		}
		if ptrB != nil {
			ptrB = ptrB.Next
		} else {
			ptrB = headA
		}
	}
	return ptrA
}
