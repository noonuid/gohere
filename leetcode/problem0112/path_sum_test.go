package problem0112

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var NULL = structure.NULL

func testFramework(t *testing.T, testFn func(*TreeNode, int) bool) {
	testCases := []struct {
		root      *TreeNode
		targetSum int
		want      bool
	}{
		{
			root:      structure.Ints2Tree([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, NULL, 1}),
			targetSum: 22,
			want:      true,
		},
		{
			root:      structure.Ints2Tree([]int{1, 2, 3}),
			targetSum: 5,
			want:      false,
		},
		{
			root:      structure.Ints2Tree([]int{}),
			targetSum: 0,
			want:      false,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFn(testCase.root, testCase.targetSum)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestPathSum(t *testing.T) {
	testFramework(t, hasPathSum)
}
