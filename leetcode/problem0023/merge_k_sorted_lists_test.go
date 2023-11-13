package problem0023

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(lists []*ListNode) *ListNode) {
	// 测试用例。
	testCases := []struct {
		lists []*ListNode
		want  []int
	}{
		{
			lists: []*ListNode{
				structure.Ints2List([]int{1, 4, 5}),
				structure.Ints2List([]int{1, 3, 4}),
				structure.Ints2List([]int{2, 6}),
			},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			lists: []*ListNode{},
			want:  []int{},
		},
		{
			lists: []*ListNode{structure.Ints2List([]int{})},
			want:  []int{},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := structure.List2Ints(testFunc(testCase.lists))
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMergeKLists_divide_and_conquer(t *testing.T) {
	testFramework(t, mergeKLists_divide_and_conquer)
}

func TestMergeKLists_sequentially(t *testing.T) {
	testFramework(t, mergeKLists_sequentially)
}
