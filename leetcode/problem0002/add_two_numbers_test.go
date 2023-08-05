package problem0002

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(*ListNode, *ListNode) *ListNode) {
	// 测试用例。
	testCases := []struct {
		l1   *ListNode
		l2   *ListNode
		want []int
	}{
		{
			l1:   structure.Ints2List([]int{2, 4, 3}),
			l2:   structure.Ints2List([]int{5, 6, 4}),
			want: []int{7, 0, 8},
		},
		{
			l1:   structure.Ints2List([]int{0}),
			l2:   structure.Ints2List([]int{0}),
			want: []int{0},
		},
		{
			l1:   structure.Ints2List([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:   structure.Ints2List([]int{9, 9, 9, 9}),
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := structure.List2Ints(testFunc(testCase.l1, testCase.l2))
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	testFramework(t, addTwoNumbers)
}
