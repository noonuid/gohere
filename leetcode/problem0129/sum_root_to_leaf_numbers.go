package problem0129

import "github.com/noonuid/go/leetcode/structure"

// 129. 求根节点到叶节点数字之和

// 给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
// 每条从根节点到叶节点的路径都代表一个数字：

// 例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
// 计算从根节点到叶节点生成的 所有数字之和 。

// 叶节点 是指没有子节点的节点。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func sumNumbers(root *TreeNode) int {
	result := 0

	var preorder func(node *TreeNode, num int)
	preorder = func(node *TreeNode, num int) {
		if node.Left == nil && node.Right == nil {
			result += num + node.Val
			return
		}
		if node.Left != nil {
			preorder(node.Left, (num+node.Val)*10)
		}
		if node.Right != nil {
			preorder(node.Right, (num+node.Val)*10)
		}

	}

	preorder(root, 0)
	return result
}
