package problem0129

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func test(t *testing.T, fn func(root *TreeNode) int) {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: structure.Ints2Tree([]int{1, 2, 3}),
			want: 25,
		},
		{
			root: structure.Ints2Tree([]int{4, 9, 0, 5, 1}),
			want: 1026,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestSumNumbers(t *testing.T) {
	test(t, sumNumbers)
}
