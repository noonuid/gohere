package problem0173

import "github.com/noonuid/gohere/leetcode/structure"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structure.TreeNode

type BSTIterator struct {
	cur   *TreeNode
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	node := root
	stack := []*TreeNode{node}
	for node.Left != nil {
		node = node.Left
		stack = append(stack, node)
	}
	return BSTIterator{
		cur:   &TreeNode{Val: -1},
		stack: stack,
	}
}

func (this *BSTIterator) Next() int {
	if node := this.cur.Right; node != nil {
		for node.Left != nil {
			this.stack = append(this.stack, node)
			node = node.Left
		}
		this.cur = node
	} else if len(this.stack) > 0 {
		this.cur = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	return this.cur.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.cur.Right != nil || len(this.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
