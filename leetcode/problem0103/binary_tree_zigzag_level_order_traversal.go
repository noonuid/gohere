package problem0103

import "github.com/noonuid/go/leetcode/structure"

// 103. 二叉树的锯齿形层序遍历

// 给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	levels := [][]int{}
	for len(stack) > 0 {
		nextStack := []*TreeNode{}
		level := []int{}
		leftToRight := len(levels)%2 == 0
		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if leftToRight {
				if node.Left != nil {
					nextStack = append(nextStack, node.Left)
				}
				if node.Right != nil {
					nextStack = append(nextStack, node.Right)
				}
			} else {
				if node.Right != nil {
					nextStack = append(nextStack, node.Right)
				}
				if node.Left != nil {
					nextStack = append(nextStack, node.Left)
				}
			}
			level = append(level, node.Val)
		}
		levels = append(levels, level)
		stack = nextStack
	}
	return levels
}
