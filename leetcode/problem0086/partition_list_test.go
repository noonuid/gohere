package problem0086

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func test(t *testing.T, fn func(head *ListNode, x int) *ListNode) {
	testCases := []struct {
		head *ListNode
		x    int
		want []int
	}{
		{
			head: structure.Ints2List([]int{1, 4, 3, 2, 5, 2}),
			x:    3,
			want: []int{1, 2, 2, 4, 3, 5},
		},
		{
			head: structure.Ints2List([]int{1, 1}),
			x:    2,
			want: []int{1, 1},
		},
	}

	for caseIndex, testCase := range testCases {
		got := structure.List2Ints(fn(testCase.head, testCase.x))
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestPartition(t *testing.T) {
	test(t, partition)
}

func TestPartition_insertion(t *testing.T) {
	test(t, partition_insertion)
}
