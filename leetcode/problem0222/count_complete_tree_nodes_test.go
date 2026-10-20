package problem0222

import (
	"reflect"
	"testing"

	"github.com/noonuid/gohere/leetcode/structure"
)

func test(t *testing.T, fn func(root *TreeNode) int) {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: structure.Ints2Tree([]int{1, 2, 3, 4, 5, 6}),
			want: 6,
		},
		{
			root: structure.Ints2Tree([]int{}),
			want: 0,
		},
		{
			root: structure.Ints2Tree([]int{1}),
			want: 1,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.root)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestCountNodes(t *testing.T) {
	test(t, countNodes)
}
