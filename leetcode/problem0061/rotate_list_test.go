package problem0061

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func test(t *testing.T, fn func(head *ListNode, k int) *ListNode) {
	testCases := []struct {
		head *ListNode
		k    int
		want []int
	}{
		{
			head: structure.Ints2List([]int{1, 2, 3, 4, 5}),
			k:    2,
			want: []int{4, 5, 1, 2, 3},
		},
		{
			head: structure.Ints2List([]int{0, 1, 2}),
			k:    4,
			want: []int{2, 0, 1},
		},
		{
			head: structure.Ints2List([]int{1}),
			k:    0,
			want: []int{1},
		},
		{
			head: structure.Ints2List([]int{1}),
			k:    1,
			want: []int{1},
		},
	}

	for caseIndex, testCase := range testCases {
		head := fn(testCase.head, testCase.k)
		got := structure.List2Ints(head)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestDeleteRotateRight(t *testing.T) {
	test(t, rotateRight)
}
