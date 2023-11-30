package problem0337

import "github.com/noonuid/go/leetcode/structure"

// 337. 打家劫舍 III

// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

// 动态规划。
func rob(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var postorder func(node *TreeNode) [2]int
	postorder = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{}
		}
		lAmounts := postorder(node.Left)
		rAmounts := postorder(node.Right)
		// 打劫当前节点能够盗取的最大金额。
		amountYes := lAmounts[1] + rAmounts[1] + node.Val
		// 不打劫当前节点能够盗取的最大金额。
		amountNo := max(lAmounts[0], lAmounts[1]) + max(rAmounts[0], rAmounts[1])
		return [2]int{amountYes, amountNo}
	}
	amounts := postorder(root)
	return max(amounts[0], amounts[1])
}
