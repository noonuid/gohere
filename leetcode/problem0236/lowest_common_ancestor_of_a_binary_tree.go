package problem0236

import (
	"github.com/noonuid/go/leetcode/structure"
)

// 236. 二叉树的最近公共祖先

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// 使用哈希表存储每个节点的父节点，再分别以 p，q 节点为起点向上遍历，找到最近的交点。
func lowestCommonAncestor_map(root, p, q *TreeNode) *TreeNode {
	parent := make(map[*TreeNode]*TreeNode)
	visited := make(map[*TreeNode]bool)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left] = node
			preorder(node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			preorder(node.Right)
		}
	}
	preorder(root)
	for node := p; node != nil; node = parent[node] {
		visited[node] = true
	}
	for node := q; node != nil; node = parent[node] {
		if visited[node] {
			return node
		}
	}
	return nil
}

// 找到从根节点到 p，q 节点的路径，两条路径公共部分的最后一个节点就是最近公共祖先。
func lowestCommonAncestor_path(root, p, q *TreeNode) *TreeNode {
	pPath, qPath := make([]*TreeNode, 0), make([]*TreeNode, 0)
	var getPath func(node, target *TreeNode, path *[]*TreeNode)
	getPath = func(node, target *TreeNode, path *[]*TreeNode) {
		*path = append(*path, node)
		if node == target {
			return
		} else {
			if node.Left != nil {
				getPath(node.Left, target, path)

			}
			if node.Right != nil {
				getPath(node.Right, target, path)
			}
		}
		if (*path)[len(*path)-1] != target {
			*path = (*path)[:len(*path)-1]
		}
	}
	getPath(root, p, &pPath)
	getPath(root, q, &qPath)
	m, n := len(pPath), len(qPath)
	i := 0
	for i < m && i < n && pPath[i] == qPath[i] {
		i++
	}
	if i > 0 {
		return pPath[i-1]
	} else {
		return nil
	}
}
