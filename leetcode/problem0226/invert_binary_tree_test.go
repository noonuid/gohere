package problem0226

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(*TreeNode) *TreeNode) {
	// 测试用例。
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: structure.Ints2Tree([]int{4, 2, 7, 1, 3, 6, 9}),
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			root: structure.Ints2Tree([]int{2, 1, 3}),
			want: []int{2, 3, 1},
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
		if !reflect.DeepEqual(structure.Tree2Ints(got), testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestInvertTree(t *testing.T) {
	testFramework(t, invertTree)
}
