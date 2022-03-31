package problem0102

import (
	"github.com/anchorportal/go/leetcode/structure"
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
func levelOrder(root *TreeNode) [][]int {
	var route = [][]int{}
	var queue = []*TreeNode{}

	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		var levelRoute = []int{}

		var length = len(queue)
		for i := 0; i < length; i++ {
			levelRoute = append(levelRoute, queue[0].Val)

			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			queue = queue[1:]
		}
		route = append(route, levelRoute)
	}

	return route
}
