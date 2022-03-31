package binarysearchtree

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(val int) *Node {
	return &Node{data: val}
}
