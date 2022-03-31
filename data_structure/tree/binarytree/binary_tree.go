package binarytree

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// 用于在整数切片中表示空的树节点
var Null = -1 << 63

// 将整型切片转换为二叉树，并返回二叉树的根节点
func IntsToTree(nums []int) *TreeNode {
	var n = len(nums)
	if n == 0 {
		return nil
	}

	var root = &TreeNode{value: nums[0]}
	var queue = make([]*TreeNode, 1, len(nums))
	queue[0] = root

	for i := 1; i < n; {
		var node = queue[0]
		queue = queue[1:]

		if i < n && nums[i] != Null {
			node.left = &TreeNode{value: nums[i]}
			queue = append(queue, node.left)
		}
		i++

		if i < n && nums[i] != Null {
			node.right = &TreeNode{value: nums[i]}
			queue = append(queue, node.right)
		}
		i++
	}

	return root
}

// 将二叉树转换为整型切片，切片中为二叉树的层序遍历结果
func TreeToInts(root *TreeNode) []int {
	return LevelOrderTraversal(root)
}

// 使用循环层序遍历二叉树
func LevelOrderTraversal(root *TreeNode) []int {
	var route = []int{}
	if root == nil {
		return route
	}

	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var node = queue[0]
		queue = queue[1:]

		route = append(route, node.value)

		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	return route
}

// 使用递归先序遍历二叉树
func PreorderTraversal(root *TreeNode) []int {
	var route = make([]int, 0)
	preorder(root, &route)
	return route
}

func preorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}

	*route = append(*route, node.value)
	preorder(node.left, route)
	preorder(node.right, route)
}

// 使用循环先序遍历二叉树
func PreorderTraversalIterative(root *TreeNode) []int {
	var route = make([]int, 0)
	var stack = make([]*TreeNode, 0)

	var node = root
	for node != nil || len(stack) > 0 {
		for node != nil {
			route = append(route, node.value)
			stack = append(stack, node)

			node = node.left
		}

		node = stack[len(stack)-1].right
		stack = stack[:len(stack)-1]
	}

	return route
}

// 使用递归中序遍历二叉树
func InorderTraversal(root *TreeNode) []int {
	var route = make([]int, 0)
	inorder(root, &route)
	return route
}

func inorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}

	inorder(node.left, route)
	*route = append(*route, node.value)
	inorder(node.right, route)
}

// 使用循环中序遍历二叉树
func InorderTraversalIterative(root *TreeNode) []int {
	var route = make([]int, 0)
	var stack = make([]*TreeNode, 0)

	var node = root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)

			node = node.left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		route = append(route, node.value)

		node = node.right
	}

	return route
}

// 使用递归后序遍历二叉树
func PostorderTraversal(root *TreeNode) []int {
	var route = make([]int, 0)
	postorder(root, &route)
	return route
}

func postorder(node *TreeNode, route *[]int) {
	if node == nil {
		return
	}

	postorder(node.left, route)
	postorder(node.right, route)
	*route = append(*route, node.value)
}

// 使用循后序遍历二叉树
func PostorderTraversalIterative(root *TreeNode) []int {
	var route = make([]int, 0)
	var stack = make([]*TreeNode, 0)

	var node, prev *TreeNode = root, nil
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)

			node = node.left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 右子树存在 && 没被访问过
		if node.right != nil && node.right != prev {
			// 再次将当前节点入栈，并开始访问其右子树
			stack = append(stack, node)
			node = node.right
		} else {
			route = append(route, node.value)
			// 避免重复访问右子树
			prev = node
			// 避免循环访问当前节点
			node = nil
		}
	}

	return route
}
