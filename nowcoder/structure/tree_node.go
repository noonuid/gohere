package structure

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// T2s convert *TreeNode to []int
func T2s(head *TreeNode, array *[]int) {
	fmt.Printf("运行到这里了 head = %v array = %v\n", head, array)
	// fmt.Printf("****array = %v\n", array)
	// fmt.Printf("****head = %v\n", head.Val)
	*array = append(*array, head.Val)
	if head.Left != nil {
		T2s(head.Left, array)
	}
	if head.Right != nil {
		T2s(head.Right, array)
	}
}
