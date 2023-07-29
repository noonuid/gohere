package problem0141

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(*ListNode) bool) {
	// 测试用例
	testCases := []struct {
		head *ListNode
		pos  int
		want bool
	}{
		{
			head: structure.Ints2List([]int{3, 2, 0, -4}),
			pos:  1,
			want: true,
		},
		{
			head: structure.Ints2List([]int{1, 2}),
			pos:  0,
			want: true,
		},
		{
			head: structure.Ints2List([]int{1}),
			pos:  -1,
			want: false,
		},
	}

	for caseIndex, testCase := range testCases {
		// 根据 pos 构造环。
		tail := testCase.head
		var tailNext *ListNode
		index := 0
		for tail != nil && tail.Next != nil {
			if index == testCase.pos {
				tailNext = tail
			}
			tail = tail.Next
			index++
		}
		if testCase.pos != -1 && tailNext != nil {
			tail.Next = tailNext
		}

		// 被测方法的返回值
		got := testFunc(testCase.head)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestHasCycle(t *testing.T) {
	testFramework(t, hasCycle)
}

func TestHasCycle_hash(t *testing.T) {
	testFramework(t, hasCycle_hash)
}
