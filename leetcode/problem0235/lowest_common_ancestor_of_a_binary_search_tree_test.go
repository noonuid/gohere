package problem0235

import (
	"reflect"
	"testing"

	"github.com/noonuid/gohere/leetcode/structure"
)

var null = structure.Null

func test(t *testing.T, fn func(root, p, q *TreeNode) *TreeNode) {
	testCases := []struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			root: structure.Ints2Tree([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}),
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 8},
			want: &TreeNode{Val: 6},
		},
		{
			root: structure.Ints2Tree([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}),
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 4},
			want: &TreeNode{Val: 2},
		},
	}

	for i := 0; i < len(testCases); i++ {
		var node *TreeNode

		node = testCases[i].root
		for node.Val != testCases[i].p.Val {
			if testCases[i].p.Val < node.Val {
				node = node.Left
			} else {
				node = node.Right
			}
		}
		testCases[i].p = node

		node = testCases[i].root
		for node.Val != testCases[i].q.Val {
			if testCases[i].q.Val < node.Val {
				node = node.Left
			} else {
				node = node.Right
			}
		}
		testCases[i].q = node

		node = testCases[i].root
		for node.Val != testCases[i].want.Val {
			if testCases[i].want.Val < node.Val {
				node = node.Left
			} else {
				node = node.Right
			}
		}
		testCases[i].want = node
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.root, testCase.p, testCase.q)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestLowestCommonAncestor(t *testing.T) {
	test(t, lowestCommonAncestor)
}
