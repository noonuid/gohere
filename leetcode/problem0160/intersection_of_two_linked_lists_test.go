package problem0160

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(*ListNode, *ListNode) *ListNode) {
	// 测试用例
	testCases := []struct {
		// 用于创建链式数据结构的输入
		intersectVal int
		listA        []int
		listB        []int
		skipA        int
		skipB        int

		// 输入数组转换后的链表的头节点指针
		headA *ListNode
		headB *ListNode

		// 预期输出
		want *ListNode
	}{
		{
			intersectVal: 8,
			listA:        []int{4, 1, 8, 4, 5},
			listB:        []int{5, 6, 1, 8, 4, 5},
			skipA:        2,
			skipB:        3,
		},
		{
			intersectVal: 2,
			listA:        []int{1, 9, 1, 2, 4},
			listB:        []int{3, 2, 4},
			skipA:        3,
			skipB:        1,
		},
		{
			intersectVal: 0,
			listA:        []int{2, 6, 4},
			listB:        []int{1, 5},
			skipA:        3,
			skipB:        2,
		},
	}

	for caseIndex, testCase := range testCases {
		// 生成交叉链表
		testCase.headA = structure.Ints2List(testCase.listA)
		testCase.headB = structure.Ints2List(testCase.listB)
		if testCase.intersectVal == 0 {
			testCase.want = nil
		} else {
			ptrA := testCase.headA
			ptrB := testCase.headB
			for countA := testCase.skipA; countA > 0; countA-- {
				ptrA = ptrA.Next
			}
			for countB := testCase.skipB; countB > 1; countB-- {
				ptrB = ptrB.Next
			}
			ptrB.Next = ptrA
			testCase.want = ptrA
		}

		got := testFunc(testCase.headA, testCase.headB)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestGetIntersectionNode(t *testing.T) {
	testFramework(t, getIntersectionNode)
}

func TestGetIntersectionNodeMap(t *testing.T) {
	testFramework(t, getIntersectionNodeMap)
}

func TestGetIntersectionNodeAlign(t *testing.T) {
	testFramework(t, getIntersectionNodeAlign)
}

func TestGetIntersectionNodeJoin(t *testing.T) {
	testFramework(t, getIntersectionNodeJoin)
}
