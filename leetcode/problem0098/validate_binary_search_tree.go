package problem0098

import (
	"math"

	"github.com/noonuid/go/leetcode/structure"
)

// 98. 验证二叉搜索树

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

// 有效 二叉搜索树定义如下：

// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func recurse(node *TreeNode, lower, upper int) bool {
	if node == nil {
		return true
	}
	if !(lower < node.Val && node.Val < upper) {
		return false
	}
	return recurse(node.Left, lower, node.Val) &&
		recurse(node.Right, node.Val, upper)
}

// 递归。
func isValidBST_recursion(root *TreeNode) bool {
	return recurse(root, math.MinInt, math.MaxInt)
}

// 迭代。
func isValidBST_iteration(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	pre := math.MinInt
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre >= node.Val {
			return false
		}
		pre = node.Val
		node = node.Right
	}
	return true
}
