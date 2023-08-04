package problem0437

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var NULL = structure.NULL

func testFramework(t *testing.T, testFunc func(*TreeNode, int) int) {
	// 测试用例
	testCases := []struct {
		root      *TreeNode
		targetSum int
		want      int
	}{
		{
			root:      structure.Ints2Tree([]int{10, 5, -3, 3, 2, NULL, 11, 3, -2, NULL, 1}),
			targetSum: 8,
			want:      3,
		},
		{
			root:      structure.Ints2Tree([]int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, 5, 1}),
			targetSum: 22,
			want:      3,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.root, testCase.targetSum)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestPathSum_prefix_sum(t *testing.T) {
	testFramework(t, pathSum_prefix_sum)
}

func TestPathSum_brute(t *testing.T) {
	testFramework(t, pathSum_brute)
}
