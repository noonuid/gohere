package problem0101

import (
	"github.com/noonuid/go/nowcoder/structure"
)

// 101. 对称二叉树

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

// 迭代移动双指针。
func isSymmetric_iteration(root *TreeNode) bool {
	p, q := root, root
	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		p, q = queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left, q.Right)
		queue = append(queue, p.Right, q.Left)
	}
	return true
}

// 递归移动双指针。
func isSymmetric_recursion(root *TreeNode) bool {
	var check func(p, q *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return check(root, root)
}
