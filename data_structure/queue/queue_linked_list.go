package queue

type QueueLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// 初始化一个新的使用链表实现的队列
func NewQueueLinkedList() *QueueLinkedList {
	return &QueueLinkedList{}
}

// 将新的元素插入到队尾
func (queue *QueueLinkedList) EnQueue(val interface{}) {
	node := NewNode(val)

	if queue.tail != nil {
		queue.tail.next = node

	}
	queue.tail = node

	if queue.head == nil {
		queue.head = node
	}

	queue.length++
}

// 删除队头元素，并返回删除的队头元素
func (queue *QueueLinkedList) DeQueue() interface{} {
	if queue.length == 0 {
		return nil
	}

	deQueued := queue.head.data
	queue.head = queue.head.next

	if queue.length == 1 {
		queue.tail = nil
	}

	queue.length--

	return deQueued
}

// 若 queue 为为空队列，则返回 TRUE，否则返回 FALSE
func (queue *QueueLinkedList) IsEmpty() bool {
	return queue.length == 0
}

// 获取队列长度
func (queue *QueueLinkedList) Length() int {
	return queue.length
}
