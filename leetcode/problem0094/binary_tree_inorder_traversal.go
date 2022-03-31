package problem0094

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

//  法一：使用递归完成
func inorderTraversal(root *TreeNode) []int {
	var route = make([]int, 0)
	inorder(root, &route)
	return route
}

func inorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, route)
	*route = append(*route, node.Val)
	inorder(node.Right, route)
}

// 法二：使用迭代完成
func inorderTraversalIterative(root *TreeNode) []int {
	var route = []int{}
	var stack = []*TreeNode{}

	var node = root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		route = append(route, node.Val)
		stack = stack[:len(stack)-1]

		node = node.Right
	}

	return route
}
