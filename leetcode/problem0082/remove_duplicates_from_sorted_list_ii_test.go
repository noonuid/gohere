package problem0082

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func test(t *testing.T, fn func(head *ListNode) *ListNode) {
	testCases := []struct {
		head *ListNode
		want []int
	}{
		{
			head: structure.Ints2List([]int{1, 2, 3, 3, 4, 4, 5}),
			want: []int{1, 2, 5},
		},
		{
			head: structure.Ints2List([]int{1, 1, 1, 2, 3}),
			want: []int{2, 3},
		},
	}

	for caseIndex, testCase := range testCases {
		head := fn(testCase.head)
		got := structure.List2Ints(head)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestDeleteDuplicates(t *testing.T) {
	test(t, deleteDuplicates)
}
