package problem0124

import (
	"math"

	"github.com/noonuid/go/leetcode/structure"
)

// 124. 二叉树中的最大路径和

// 二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

// 路径和 是路径中各节点值的总和。

// 给你一个二叉树的根节点 root ，返回其 最大路径和 。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

// 递归。
func maxPathSum(root *TreeNode) int {
	pathSum := math.MinInt
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		pathSum = max(leftGain+node.Val+rightGain, pathSum)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return pathSum
}
