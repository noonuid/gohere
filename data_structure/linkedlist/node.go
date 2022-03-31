package linkedlist

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func NewNode(val interface{}) *Node {
	return &Node{val, nil, nil}
}
