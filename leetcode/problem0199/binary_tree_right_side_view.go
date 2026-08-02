package problem0199

import "github.com/noonuid/go/leetcode/structure"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		result = append(result, queue[0].Val)
		nextQueue := []*TreeNode{}
		for _, node := range queue {
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
		}
		queue = nextQueue
	}
	return result
}

func rightSideView_dfs(root *TreeNode) []int {
	result := []int{}
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > len(result) {
			result = append(result, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 1)
	return result
}
