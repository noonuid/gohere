package problem0024

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(head *ListNode) *ListNode) {
	// 测试用例。
	testCases := []struct {
		head *ListNode
		want []int
	}{
		{
			head: structure.Ints2List([]int{1, 2, 3, 4}),
			want: []int{2, 1, 4, 3},
		},
		{
			head: structure.Ints2List([]int{}),
			want: []int{},
		},
		{
			head: structure.Ints2List([]int{1}),
			want: []int{1},
		},
		{
			head: structure.Ints2List([]int{1, 2, 3}),
			want: []int{2, 1, 3},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		newHead := testFunc(testCase.head)
		got := structure.List2Ints(newHead)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestSwapPairs(t *testing.T) {
	testFramework(t, swapPairs)
}

func TestSwapPairs_iteration(t *testing.T) {
	testFramework(t, swapPairs_iteration)
}
