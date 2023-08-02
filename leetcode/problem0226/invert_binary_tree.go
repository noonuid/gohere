package problem0226

import "github.com/noonuid/go/leetcode/structure"

// 226. 翻转二叉树

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
