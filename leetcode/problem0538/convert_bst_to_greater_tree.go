package problem0538

import "github.com/noonuid/go/leetcode/structure"

// 538. 把二叉搜索树转换为累加树

// 给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和。

// 提醒一下，二叉搜索树满足下列约束条件：

// 节点的左子树仅包含键 小于 节点键的节点。
// 节点的右子树仅包含键 大于 节点键的节点。
// 左右子树也必须是二叉搜索树。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

// 使用 Morris 方法反向中序遍历，空间复杂度为 O(1)。
func convertBST_morris(root *TreeNode) *TreeNode {
	var getPrecursor = func(node *TreeNode) *TreeNode {
		pre := node.Right
		for pre.Left != nil && pre.Left != node {
			pre = pre.Left
		}
		return pre
	}
	sum := 0
	node := root
	for node != nil {
		if node.Right == nil {
			sum += node.Val
			node.Val = sum
			node = node.Left
		} else {
			// 如果当前节点的右子节点不为空，则获取当前节点的前驱节点。
			// 初次获取时，将前驱节点的左指针指向当前节点，然后遍历当前节点的右子树。
			// 二次获取时，将前驱节点的左指针还原置空，并处理当前节点，然后遍历当前节点的左子树。
			pre := getPrecursor(node)
			if pre.Left == nil {
				pre.Left = node
				node = node.Right
			} else {
				pre.Left = nil
				sum += node.Val
				node.Val = sum
				node = node.Left
			}
		}
	}
	return root
}

// 反向中序遍历。
func convertBST_reverse_inorder(root *TreeNode) *TreeNode {
	sum := 0
	var reverseInrder func(node *TreeNode)
	reverseInrder = func(node *TreeNode) {
		if node != nil {
			reverseInrder(node.Right)
			sum += node.Val
			node.Val = sum
			reverseInrder(node.Left)
		}
	}
	reverseInrder(root)
	return root
}
