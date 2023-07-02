package structure

import "fmt"

// NULL 方便添加测试数据。
var NULL = -1 << 63

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Ints2Tree 根据 []int 生成树，并返回树的根节点。
// 借助队列实现。
func Ints2Tree(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n)
	queue[0] = root

	for i := 1; i < n && len(queue) > 0; {
		node := queue[0]
		queue = queue[1:]

		if ints[i] != NULL {
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

// Tree2Ints converts a tree to slice.
// Use NULL to represent the nil.
func Tree2Ints(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, node.Val)

			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			res = append(res, NULL)
		}
	}
	for len(res) > 0 && res[len(res)-1] == NULL {
		res = res[:len(res)-1]
	}

	return res
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

// 先序遍历。
func PreorderTraversal(root *TreeNode) []int {
	var route = make([]int, 0)
	preorder(root, &route)
	return route
}

func preorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}

	*route = append(*route, node.Val)
	preorder(node.Left, route)
	preorder(node.Right, route)
}

// 中序遍历。
func InorderTraversal(root *TreeNode) []int {
	var route = make([]int, 0)
	inorder(root, &route)
	return route
}

func inorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, route)
	*route = append(*route, node.Val)
	inorder(node.Right, route)
}

// 后序遍历。
func PostorderTraversal(root *TreeNode) []int {
	var route = make([]int, 0)
	postorder(root, &route)
	return route
}

func postorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}

	postorder(node.Left, route)
	postorder(node.Right, route)
	*route = append(*route, node.Val)
}
