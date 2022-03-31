package queue

type Node struct {
	data interface{}
	next *Node
}

func NewNode(val interface{}) *Node {
	return &Node{data: val}
}
