package problem0092

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(head *ListNode, left int, right int) *ListNode) {
	testCases := []struct {
		head  *ListNode
		left  int
		right int
		want  []int
	}{
		{
			head:  structure.Ints2List([]int{1, 2, 3, 4, 5}),
			left:  2,
			right: 4,
			want:  []int{1, 4, 3, 2, 5},
		},
		{
			head:  structure.Ints2List([]int{5}),
			left:  1,
			right: 1,
			want:  []int{5},
		},
		{
			head:  structure.Ints2List([]int{3, 5}),
			left:  1,
			right: 1,
			want:  []int{3, 5},
		},
	}

	for caseIndex, testCase := range testCases {
		head := testFunc(testCase.head, testCase.left, testCase.right)
		got := structure.List2Ints(head)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestReverseBetween(t *testing.T) {
	testFramework(t, reverseBetween)
}

func TestReverseBetween_concatenate(t *testing.T) {
	testFramework(t, reverseBetween_concatenate)
}
