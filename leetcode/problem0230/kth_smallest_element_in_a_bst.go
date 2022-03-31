package problem0230

import (
	"github.com/anchorportal/go/leetcode/structure"
)

type TreeNode = structure.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var route []int
	inorder(root, &route)

	return route[k-1]
}

func inorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, route)
	*route = append(*route, node.Val)
	inorder(node.Right, route)
}
