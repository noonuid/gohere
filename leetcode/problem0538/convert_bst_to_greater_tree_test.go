package problem0538

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var null = structure.NULL

func testFramework(t *testing.T, testFunc func(root *TreeNode) *TreeNode) {
	// 测试用例。
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: structure.Ints2Tree([]int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8}),
			want: []int{30, 36, 21, 36, 35, 26, 15, null, null, null, 33, null, null, null, 8},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := structure.Tree2Ints(testFunc(testCase.root))
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestTonvertBST_morris(t *testing.T) {
	testFramework(t, convertBST_morris)
}

func TestTonvertBST_reverse_inorder(t *testing.T) {
	testFramework(t, convertBST_reverse_inorder)
}
