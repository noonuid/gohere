package problem0019

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(head *ListNode, n int) *ListNode) {
	// 测试用例。
	testCases := []struct {
		head *ListNode
		n    int
		want []int
	}{
		{
			head: structure.Ints2List([]int{1, 2, 3, 4, 5}),
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			head: structure.Ints2List([]int{1}),
			n:    1,
			want: []int{},
		},
		{
			head: structure.Ints2List([]int{1, 2}),
			n:    1,
			want: []int{1},
		},
	}

	for caseIndex, testCase := range testCases {
		got := structure.List2Ints(testFunc(testCase.head, testCase.n))
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.n)
		}
	}
}

func TestRemoveNthFromEnd_double_pointer(t *testing.T) {
	testFramework(t, removeNthFromEnd_double_pointer)
}

func TestRemoveNthFromEnd_stack(t *testing.T) {
	testFramework(t, removeNthFromEnd_stack)
}

func TestRemoveNthFromEnd_two_traversal(t *testing.T) {
	testFramework(t, removeNthFromEnd_two_traversal)
}
