package problem0617

// 给你两棵二叉树： root1 和 root2 。

// 想象一下，当你将其中一棵覆盖到另一棵之上时，两棵树上的一些节点将会重叠（而另一些不会）。你需要将这两棵树合并成一棵新二叉树。合并的规则是：如果两个节点重叠，那么将这两个节点的值相加作为合并后节点的新值；否则，不为 null 的节点将直接作为新二叉树的节点。

// 返回合并后的二叉树。

// 注意: 合并过程必须从两个树的根节点开始。

import (
	"github.com/noonuid/go/leetcode/structure"
)

type TreeNode = structure.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var root *TreeNode

	switch {
	case root1 != nil && root2 != nil:
		root = &TreeNode{
			Val: root1.Val + root2.Val,
		}
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
	case root1 != nil && root2 == nil:
		root = &TreeNode{
			Val: root1.Val,
		}
		root.Left = mergeTrees(root1.Left, nil)
		root.Right = mergeTrees(root1.Right, nil)
	case root1 == nil && root2 != nil:
		root = &TreeNode{
			Val: root2.Val,
		}
		root.Left = mergeTrees(nil, root2.Left)
		root.Right = mergeTrees(nil, root2.Right)
	default:
	}

	return root
}
