package problem0222

import (
	"sort"

	"github.com/noonuid/gohere/leetcode/structure"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	firstNilNode := sort.Search(1<<(level+1), func(i int) bool {
		if i <= 1<<level {
			return false
		}
		node := root
		mask := 1 << (level - 1)
		for node != nil && mask > 0 {
			if i&mask == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			mask = mask >> 1
		}
		return node == nil
	})
	return firstNilNode - 1
}
