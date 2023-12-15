package problem0104

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
			root: structure.Ints2Tree([]int{3, 9, 20, null, null, 15, 7}),
			want: 3,
		},
		{
			root: structure.Ints2Tree([]int{1, null, 2}),
			want: 2,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.root)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMaxDepth(t *testing.T) {
	testFramework(t, maxDepth)
}
