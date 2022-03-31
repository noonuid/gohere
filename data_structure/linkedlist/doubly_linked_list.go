package linkedlist

type DoublyLinkedList struct {
	head   *Node
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{nil, 0}
}

func (list *DoublyLinkedList) AddAtBeg(val interface{}) {
	node := NewNode(val)
	node.next = list.head

	if list.head != nil {
		list.head.prev = node
	}

	list.head = node
	list.length++
}

func (list *DoublyLinkedList) AddAtEnd(val interface{}) {
	node := NewNode(val)

	if list.head == nil {
		list.head = node
		return
	}

	var cur *Node
	for cur = list.head; cur.next != nil; cur = cur.next {
	}
	cur.next = node
	node.prev = cur
	list.length++
}

func (list *DoublyLinkedList) DelAtBeg() interface{} {
	if list.head == nil {
		return -1
	}

	cur := list.head
	list.head = cur.next
	if list.head != nil {
		list.head.prev = nil
	}

	list.length--
	return cur.data
}

func (list *DoublyLinkedList) DelAtEnd() interface{} {
	if list.head == nil {
		return -1
	}

	cur := list.head
	if cur.next == nil {
		list.head = nil
		list.length--
		return cur.data
	}

	for ; cur.next != nil; cur = cur.next {
	}

	cur.prev.next = nil
	list.length--
	return cur.data
}

func (list *DoublyLinkedList) Count() int {
	return list.length
}

func (list *DoublyLinkedList) Reverse() {
	var prev, next *Node
	cur := list.head

	for cur != nil {
		next = cur.next
		cur.next = prev
		cur.prev = next
		prev = cur
		cur = next
	}

	list.head = prev
}
