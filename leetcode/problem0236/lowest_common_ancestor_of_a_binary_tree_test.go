package problem0236

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var null = structure.NULL

func testFramework(t *testing.T, testFunc func(root, p, q *TreeNode) *TreeNode) {
	// 测试用例。
	testCases := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			root: structure.Ints2Tree([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}),
		},
		{
			root: structure.Ints2Tree([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}),
		},
		{
			root: structure.Ints2Tree([]int{1, 2}),
		},
	}

	testCases[0].p = testCases[0].root.Left
	testCases[0].q = testCases[0].root.Right
	testCases[0].want = testCases[0].root

	testCases[1].p = testCases[1].root.Left
	testCases[1].q = testCases[1].root.Left.Right.Right
	testCases[1].want = testCases[1].root.Left

	testCases[2].p = testCases[2].root
	testCases[2].q = testCases[2].root.Left
	testCases[2].want = testCases[2].root

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.root, testCase.p, testCase.q)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	testFramework(t, lowestCommonAncestor)
}

func TestLowestCommonAncestor_map(t *testing.T) {
	testFramework(t, lowestCommonAncestor_map)
}

func TestLowestCommonAncestor_path(t *testing.T) {
	testFramework(t, lowestCommonAncestor_path)
}
