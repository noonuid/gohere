package problem0543

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

func testFramework(t *testing.T, testFunc func(*TreeNode) int) {
	// 测试用例
	testCases := []struct {
		root *TreeNode
		want int
	}{
		{
			root: structure.Ints2Tree([]int{1, 2, 3, 4, 5}),
			want: 3,
		},
		{
			root: structure.Ints2Tree([]int{1, 2}),
			want: 1,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.root)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestDiameterOfBinaryTree(t *testing.T) {
	testFramework(t, diameterOfBinaryTree)
}
