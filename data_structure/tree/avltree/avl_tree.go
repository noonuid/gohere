package avltree

type AvlTree struct {
	root *Node
}

func NewAvlTree() *AvlTree {
	return &AvlTree{}
}

func llRotation(node *Node) *Node {
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node
	return newRoot
}

func lrRotation(node *Node) *Node {
	node.left = rrRotation(node.left)
	return llRotation(node)
}

func rrRotation(node *Node) *Node {
	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node
	return newRoot
}

func rlRotation(node *Node) *Node {
	node.right = llRotation(node.right)
	return rrRotation(node)
}

func balanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

func height(node *Node) int {
	if node == nil {
		return -1
	}

	return max(height(node.left), height(node.right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 插入节点
func Insert(node *Node, data int) *Node {
	if node == nil {
		return NewNode(data)
	}

	if data < node.data {
		node.left = Insert(node.left, data)
	} else {
		node.right = Insert(node.right, data)
	}

	bFactor := balanceFactor(node)
	if bFactor == 2 { // L
		bFactor = balanceFactor(node.left)
		if bFactor == 1 { // LL
			return llRotation(node)
		} else if bFactor == -1 { // LR
			return lrRotation(node)
		}
	} else if bFactor == -2 { // R
		bFactor = balanceFactor(node.right)
		if bFactor == 1 { // RL
			return rlRotation(node)
		} else if bFactor == -1 { // RR
			return rrRotation(node)
		}
	}

	return node
}

// 获取树的深度
func Depth(tree *AvlTree) int {
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
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else {
			auxiliary := MinNode(node.right)
			auxiliary.left = node.left
			node = node.right
		}
	}

	bFactor := balanceFactor(node)

	if bFactor == 2 { // L
		switch balanceFactor(node.left) {
		case 1: // LL
			return llRotation(node)
		case -1: // LR
			return lrRotation(node)
		case 0: //  same as LL
			return llRotation(node)
		}
	} else if bFactor == -2 { // L
		switch balanceFactor(node.right) {
		case 1: // RL
			return rlRotation(node)
		case -1: // RR
			return rrRotation(node)
		case 0: // same as RR
			return rrRotation(node)
		}
	}

	return node
}
