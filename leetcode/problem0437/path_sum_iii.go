package problem0437

import "github.com/noonuid/go/leetcode/structure"

// 437. 路径总和 III

// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

// 前缀和。
func pathSum_prefix_sum(root *TreeNode, targetSum int) int {
	count := 0
	preSum := map[int]int{0: 1}

	var dfs func(node *TreeNode, curSum int)
	dfs = func(node *TreeNode, curSum int) {
		if node == nil {
			return
		}
		curSum += node.Val
		count += preSum[curSum-targetSum]
		preSum[curSum]++
		dfs(node.Left, curSum)
		dfs(node.Right, curSum)
		preSum[curSum]--
	}

	dfs(root, 0)
	return count
}

// 暴力枚举。
func pathSum_brute(root *TreeNode, targetSum int) int {
	answer := 0
	if root == nil {
		return answer
	}

	var sum func(node *TreeNode, targetSum int) int
	sum = func(node *TreeNode, targetSum int) int {
		count := 0
		if node == nil {
			return count
		}
		if node.Val == targetSum {
			count += 1
		}
		count += sum(node.Left, targetSum-node.Val)
		count += sum(node.Right, targetSum-node.Val)
		return count
	}

	answer += sum(root, targetSum)
	answer += pathSum_brute(root.Left, targetSum)
	answer += pathSum_brute(root.Right, targetSum)
	return answer
}
