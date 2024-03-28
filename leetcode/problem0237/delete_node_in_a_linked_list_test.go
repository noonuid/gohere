package problem0237

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(node *ListNode)) {
	testCases := []struct {
		root *ListNode
		node *ListNode
		want []int
	}{
		{
			root: structure.Ints2List([]int{4, 5, 1, 9}),
			want: []int{4, 1, 9},
		},
		{
			root: structure.Ints2List([]int{4, 5, 1, 9}),
			want: []int{4, 5, 9},
		},
	}

	testCases[0].node = testCases[0].root.Next
	testCases[1].node = testCases[1].root.Next.Next

	for caseIndex, testCase := range testCases {
		testFunc(testCase.node)
		got := structure.List2Ints(testCase.root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestDeleteNode(t *testing.T) {
	testFramework(t, deleteNode)
}
