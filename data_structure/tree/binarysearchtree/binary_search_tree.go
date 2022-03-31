package binarysearchtree

type BinarySearchTree struct {
	root *Node
}

// 插入节点
func Insert(node *Node, val int) *Node {
	if node == nil {
		return NewNode(val)
	}

	if val < node.data {
		node.left = Insert(node.left, val)
	} else {
		node.right = Insert(node.right, val)
	}

	return node
}

// 获取树的深度
func Depth(tree *BinarySearchTree) int {
	return calculateDepth(tree.root, 0)
}

// 计算深度
func calculateDepth(node *Node, depth int) int {
	if node == nil {
		return depth
	}

	leftDepth := calculateDepth(node.left, depth+1)
	rightDepth := calculateDepth(node.right, depth+1)

	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}

// 先序遍历
func PreOrder(root *Node) []int {
	traversal := make([]int, 0)
	PreOrderRecursive(root, &traversal)
	return traversal
}

// 先序遍历辅助函数
func PreOrderRecursive(node *Node, traversal *[]int) {
	if node != nil {
		*traversal = append(*traversal, node.data)
		PreOrderRecursive(node.left, traversal)
		PreOrderRecursive(node.right, traversal)
	}
}

// 中序遍历
func InOrder(root *Node) []int {
	traversal := make([]int, 0)
	InOrderRecursive(root, &traversal)
	return traversal
}

// 中序遍历辅助函数
func InOrderRecursive(node *Node, traversal *[]int) {
	if node != nil {
		InOrderRecursive(node.left, traversal)
		*traversal = append(*traversal, node.data)
		InOrderRecursive(node.right, traversal)
	}
}

// 后序遍历节点
func PostOrder(root *Node) []int {
	traversal := make([]int, 0)
	PostOrderRecursive(root, &traversal)
	return traversal
}

// 后序遍历节点辅助函数
func PostOrderRecursive(node *Node, traversal *[]int) {
	if node != nil {
		PostOrderRecursive(node.left, traversal)
		PostOrderRecursive(node.right, traversal)
		*traversal = append(*traversal, node.data)
	}
}

// 返回最小节点地址
func MinNode(node *Node) *Node {
	cur := node
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// 返回最大节点地址
func MaxNode(node *Node) *Node {
	cur := node
	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

// 删除某个节点
func Remove(node *Node, val int) *Node {
	if node == nil {
		return nil
	}

	if val < node.data {
		node.left = Remove(node.left, val)
	} else if val > node.data {
		node.right = Remove(node.right, val)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		} else {
			auxiliary := MinNode(node.right)
			auxiliary.left = node.left
			return node.right
		}
	}

	return node
}
