package problem0098

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var NULL int = structure.NULL

func testFramework(t *testing.T, testFunc func(*TreeNode) bool) {
	// 测试用例
	testCases := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: structure.Ints2Tree([]int{2, 1, 3}),
			want: true,
		},
		{
			root: structure.Ints2Tree([]int{5, 1, 4, NULL, NULL, 3, 6}),
			want: false,
		},
		{
			root: structure.Ints2Tree([]int{5, 4, 6, NULL, NULL, 3, 7}),
			want: false,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.root)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestIsValidBST_recursion(t *testing.T) {
	testFramework(t, isValidBST_recursion)
}

func TestIsValidBST_iteration(t *testing.T) {
	testFramework(t, isValidBST_iteration)
}
