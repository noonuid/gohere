package problem0025

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(head *ListNode, k int) *ListNode) {
	// 测试用例。
	testCases := []struct {
		head *ListNode
		k    int
		want []int
	}{
		{
			head: structure.Ints2List([]int{1, 2, 3, 4, 5}),
			k:    2,
			want: []int{2, 1, 4, 3, 5},
		},
		{
			head: structure.Ints2List([]int{1, 2, 3, 4, 5}),
			k:    3,
			want: []int{3, 2, 1, 4, 5},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		newHead := testFunc(testCase.head, testCase.k)
		got := structure.List2Ints(newHead)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestReverseKGroup(t *testing.T) {
	testFramework(t, reverseKGroup)
}
