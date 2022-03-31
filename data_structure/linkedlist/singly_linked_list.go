package linkedlist

import "fmt"

type SinglyLinkedList struct {
	head   *Node
	length int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (list *SinglyLinkedList) AddAtBeg(val interface{}) {
	newNode := NewNode(val)
	newNode.next = list.head
	list.head = newNode
	list.length++
}

func (list *SinglyLinkedList) AddAtEnd(val interface{}) {
	newNode := NewNode(val)

	if list.head == nil {
		list.head = newNode
	} else {
		cur := list.head
		for ; cur.next != nil; cur = cur.next {
		}
		cur.next = newNode
	}

	list.length++
}

func (list *SinglyLinkedList) DelAtBeg() interface{} {
	if list.length == 0 {
		return -1
	}

	cur := list.head
	list.head = list.head.next
	list.length--

	return cur.data
}

func (list *SinglyLinkedList) DelAtEnd() interface{} {
	if list.length == 0 {
		return -1
	}

	cur := list.head
	if list.length == 1 {
		list.head = nil
		return cur.data
	}

	for ; cur.next.next != nil; cur = cur.next {
	}

	tail := cur.next
	cur.next = nil
	list.length--

	return tail.data
}

func (list *SinglyLinkedList) Count() int {
	return list.length
}

func (list *SinglyLinkedList) Reverse() {
	var prev, next *Node
	cur := list.head

	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	list.head = prev
}

func (list *SinglyLinkedList) Display() {
	for cur := list.head; cur != nil; cur = cur.next {
		fmt.Print(cur.data, " ")
	}
	fmt.Print("\n")
}
