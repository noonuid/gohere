package stack

type StackLinkedList struct {
	top    *Node
	length int
}

// 初始化一个新的链表栈
func NewStackLinkedList() *StackLinkedList {
	return &StackLinkedList{}
}

// 将一个元素入栈
func (stack *StackLinkedList) Push(val interface{}) {
	node := NewNode(val)

	node.next = stack.top
	stack.top = node
	stack.length++
}

// 将一个元素出栈
func (stack *StackLinkedList) Pop() interface{} {
	if stack.top == nil {
		return nil
	}

	val := stack.top.data
	stack.top = stack.top.next
	stack.length--

	return val
}

// 判断当前栈是否为空
func (stack *StackLinkedList) IsEmpty() bool {
	return stack.length == 0
}

// 返回当前栈的元素个数
func (stack *StackLinkedList) Length() int {
	return stack.length
}
