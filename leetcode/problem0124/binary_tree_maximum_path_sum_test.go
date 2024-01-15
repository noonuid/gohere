package problem0124

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var null = structure.NULL

func testFramework(t *testing.T, testFunc func(root *TreeNode) int) {
	// 测试用例。
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: structure.Ints2Tree([]int{1, 2, 3}),
			want: 6,
		},
		{
			root: structure.Ints2Tree([]int{-10, 9, 20, null, null, 15, 7}),
			want: 42,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ntest case: %d\ngot : %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMaxPathSum(t *testing.T) {
	testFramework(t, maxPathSum)
}
