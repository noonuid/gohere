package problem0105

import "github.com/noonuid/go/leetcode/structure"

// 105. 从前序与中序遍历序列构造二叉树
// 给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历，
// inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

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
func buildTree_recursion(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for i < len(inorder) && inorder[i] != preorder[0] {
		i++
	}
	root.Left = buildTree_recursion(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree_recursion(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// 迭代。
func buildTree_iteration(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	inorderIndex := 0
	for i := 1; i < len(preorder); i++ {
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{Val: preorder[i]}
			stack = append(stack, node.Left)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{Val: preorder[i]}
			stack = append(stack, node.Right)
		}
	}
	return root
}
