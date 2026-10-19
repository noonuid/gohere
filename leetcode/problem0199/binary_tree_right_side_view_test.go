package problem0199

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var null int = structure.Null

func testFramework(t *testing.T, testFunc func(*TreeNode) []int) {
	// 测试用例。
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: structure.Ints2Tree([]int{1, 2, 3, null, 5, null, 4}),
			want: []int{1, 3, 4},
		},
		{
			root: structure.Ints2Tree([]int{1, 2, 3, 4, null, null, null, 5}),
			want: []int{1, 3, 4, 5},
		},
		{
			root: structure.Ints2Tree([]int{1, null, 3}),
			want: []int{1, 3},
		},
		{
			root: structure.Ints2Tree([]int{}),
			want: []int{},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.root)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ntestCase: %v\ngot: %v\nwant: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestRightSideView(t *testing.T) {
	testFramework(t, rightSideView)
}

func TestRightSideView_dfs(t *testing.T) {
	testFramework(t, rightSideView_dfs)
}
