package problem0530

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var null = structure.NULL

func testFramework(t *testing.T, testFunc func(root *TreeNode) int) {
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: structure.Ints2Tree([]int{4, 2, 6, 1, 3}),
			want: 1,
		},
		{
			root: structure.Ints2Tree([]int{1, 0, 48, null, null, 12, 49}),
			want: 1,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestGetMinimumDifference(t *testing.T) {
	testFramework(t, getMinimumDifference)
}
