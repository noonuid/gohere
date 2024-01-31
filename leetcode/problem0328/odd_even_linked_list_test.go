package problem0328

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(head *ListNode) *ListNode) {
	testCases := []struct {
		head *ListNode
		want []int
	}{
		{
			head: structure.Ints2List([]int{1, 2, 3, 4, 5}),
			want: []int{1, 3, 5, 2, 4},
		},
		{
			head: structure.Ints2List([]int{2, 1, 3, 5, 6, 4, 7}),
			want: []int{2, 3, 6, 7, 1, 5, 4},
		},
	}

	for caseIndex, testCase := range testCases {
		got := structure.List2Ints(testFunc(testCase.head))
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ntest case: %d\ngot : %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestOddEvenList(t *testing.T) {
	testFramework(t, oddEvenList)
}
