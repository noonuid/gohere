package offer06

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(*ListNode) []int) {
	// 测试用例。
	testCases := []struct {
		head *ListNode
		want []int
	}{
		{
			head: structure.Ints2List([]int{1, 3, 2}),
			want: []int{2, 3, 1},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.head)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestReversePrint(t *testing.T) {
	testFramework(t, reversePrint)
}
