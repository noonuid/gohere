package nc081

import (
	"github.com/noonuid/go/nowcoder/structure"
)

type TreeNode = structure.TreeNode

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param proot TreeNode类
 * @param k int整型
 * @return int整型
 */
func KthNode(proot *TreeNode, k int) int {
	if k < 1 {
		return -1
	}

	var stack []*TreeNode
	var node = proot
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}

		node = node.Right
	}

	return -1
}
