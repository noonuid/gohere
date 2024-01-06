package problem0114

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

// 寻找前驱节点。
func flatten_predecessor(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			// 在左子树中寻找右子树的前驱节点，该节点是左子树最右边的节点。
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur.Right
			cur.Right = cur.Left
			cur.Left = nil
		}
		cur = cur.Right
	}
}

// 前序遍历和更新指针同时进行。
func flatten_traverse_with_update(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	pre := &TreeNode{}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		pre.Left, pre.Right = nil, cur
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		pre = cur
	}
}

// 先完成前序遍历，再按顺序更改所有节点的左右指针。
func flatten_traverse_first(root *TreeNode) {
	path, stack := []*TreeNode{}, []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			path = append(path, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	for i := 0; i < len(path)-1; i++ {
		path[i].Left, path[i].Right = nil, path[i+1]
	}
}
