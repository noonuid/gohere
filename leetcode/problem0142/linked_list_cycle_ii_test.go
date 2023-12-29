package problem0142

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(head *ListNode) *ListNode) {
	// 测试用例。
	testCases := []struct {
		head *ListNode
		pos  int
		want *ListNode
	}{
		{
			head: structure.Ints2List([]int{3, 2, 0, -4}),
			pos:  1,
		},
		{
			head: structure.Ints2List([]int{1, 2}),
			pos:  0,
		},
		{
			head: structure.Ints2List([]int{1}),
			pos:  -1,
		},
	}

	for caseIndex, testCase := range testCases {
		// 根据 pos 设置入环节点和环。
		if testCase.pos != -1 {
			node := testCase.head
			for i := 0; i < testCase.pos && node != nil; i++ {
				node = node.Next
			}
			testCase.want = node
			for node.Next != nil {
				node = node.Next
			}
			node.Next = testCase.want
		}

		got := testFunc(testCase.head)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestDetectCycle_double_pointer(t *testing.T) {
	testFramework(t, detectCycle_double_pointer)
}

func TestDetectCycle_hash(t *testing.T) {
	testFramework(t, detectCycle_hash)
}
