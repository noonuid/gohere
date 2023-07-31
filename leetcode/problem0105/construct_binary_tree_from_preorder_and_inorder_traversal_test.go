package problem0105

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var NULL = structure.NULL

func testFramework(t *testing.T, testFunc func([]int, []int) *TreeNode) {
	// 测试用例
	testCases := []struct {
		preorder []int
		inorder  []int
		want     []int
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     []int{3, 9, 20, NULL, NULL, 15, 7},
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     []int{-1},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.preorder, testCase.inorder)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(structure.Tree2Ints(got), testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestBuildTree_recursion(t *testing.T) {
	testFramework(t, buildTree_recursion)
}

func TestBuildTree_iteration(t *testing.T) {
	testFramework(t, buildTree_iteration)
}
