package problem0144

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
func preorderTraversal(root *TreeNode) []int {
	var route = []int{}
	preorder(root, &route)
	return route
}

func preorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}

	*route = append(*route, node.Val)
	preorder(node.Left, route)
	preorder(node.Right, route)
}

// 法二：使用迭代完成
func preorderTraversalIterative(root *TreeNode) []int {
	var route = make([]int, 0)
	var stack = make([]*TreeNode, 0)

	var node = root
	for node != nil || len(stack) > 0 {
		for node != nil {
			route = append(route, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return route
}
