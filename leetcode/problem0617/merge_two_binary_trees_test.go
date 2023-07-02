package problem0617

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var NULL = structure.NULL

func testFramework(t *testing.T, testFunc func(root1 *TreeNode, root2 *TreeNode) *TreeNode) {
	// 测试用例
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			nums1: []int{1, 3, 2, 5},
			nums2: []int{2, 1, 3, NULL, 4, NULL, 7},
			want:  []int{3, 4, 5, 5, 4, NULL, 7},
		},
		{
			nums1: []int{1},
			nums2: []int{1, 2},
			want:  []int{2, 2},
		},
	}

	for caseIndex, testCase := range testCases {
		root1 := structure.Ints2Tree(testCase.nums1)
		root2 := structure.Ints2Tree(testCase.nums2)
		root := testFunc(root1, root2)
		got := structure.Tree2Ints(root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestMergeTrees(t *testing.T) {
	testFramework(t, mergeTrees)
}
