package problem0530

import (
	"math"

	"github.com/noonuid/go/leetcode/structure"
)

// 530. 二叉搜索树的最小绝对差

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

// 差值是一个正数，其数值等于两值之差的绝对值。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func getMinimumDifference(root *TreeNode) int {
	minDif := math.MaxInt32
	pre := math.MinInt32
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if curDif := node.Val - pre; curDif < minDif {
			minDif = curDif
		}
		pre = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return minDif
}
