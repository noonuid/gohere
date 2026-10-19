package problem0117

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		nextQueue := []*Node{}
		var pre *Node
		for _, cur := range queue {
			if pre != nil {
				pre.Next = cur
			}
			if cur.Left != nil {
				nextQueue = append(nextQueue, cur.Left)
			}
			if cur.Right != nil {
				nextQueue = append(nextQueue, cur.Right)
			}
			pre = cur
		}
		queue = nextQueue
	}
	return root
}

func connect_use_next(root *Node) *Node {
	head := root
	for head != nil {
		var nextHead, pre *Node
		for node := head; node != nil; node = node.Next {
			if cur := node.Left; cur != nil {
				if nextHead == nil {
					nextHead = cur
				} else {
					pre.Next = cur
				}
				pre = cur
			}
			if cur := node.Right; cur != nil {
				if nextHead == nil {
					nextHead = cur
				} else {
					pre.Next = cur
				}
				pre = cur
			}
		}
		head = nextHead
	}
	return root
}
