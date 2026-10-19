package problem0106

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	index := 0
	for index < len(inorder) && inorder[index] != root.Val {
		index++
	}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func buildTree_iteration(inorder []int, postorder []int) *TreeNode {
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	index := len(inorder) - 1
	for cur := len(postorder) - 2; cur > -1; cur-- {
		node := &TreeNode{Val: postorder[cur]}
		top := stack[len(stack)-1]
		if top.Val != inorder[index] {
			top.Right = node
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[index] {
				top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				index--
			}
			top.Left = node
		}
		stack = append(stack, node)
	}
	return root
}
