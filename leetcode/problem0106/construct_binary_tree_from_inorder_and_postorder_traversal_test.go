package problem0106

import (
	"reflect"
	"testing"

	"github.com/noonuid/gohere/leetcode/structure"
)

var null = structure.Null

func testFramework(t *testing.T, testFunc func(inorder []int, postorder []int) *TreeNode) {
	testCases := []struct {
		inorder   []int
		postorder []int
		want      []int
	}{
		{
			inorder:   []int{9, 3, 15, 20, 7},
			postorder: []int{9, 15, 7, 20, 3},
			want:      []int{3, 9, 20, null, null, 15, 7},
		},
		{
			inorder:   []int{-1},
			postorder: []int{-1},
			want:      []int{-1},
		},
	}

	for caseIndex, testCase := range testCases {
		root := testFunc(testCase.inorder, testCase.postorder)
		got := structure.Tree2Ints(root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestBuildTree_recursion(t *testing.T) {
	testFramework(t, buildTree)
}

func TestBuildTree_iteration(t *testing.T) {
	testFramework(t, buildTree_iteration)
}
