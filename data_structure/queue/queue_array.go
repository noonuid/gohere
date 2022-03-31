package queue

type QueueArray struct {
	data   []interface{}
	length int
}

// 初始化一个空的新队列
func NewQueueArray() *QueueArray {
	return &QueueArray{}
}

// 将一个元素插入到队尾
func (queue *QueueArray) EnQueue(val interface{}) {
	queue.data = append(queue.data, val)
	queue.length++
}

// 将队头元素出队，并返回该元素
func (queue *QueueArray) DeQueue() interface{} {
	if queue.length == 0 {
		return nil
	}

	deQueued := queue.data[0]

	queue.data = queue.data[1:]
	queue.length--

	return deQueued
}

// 获取队列长度
func (queue *QueueArray) Length() int {
	return queue.length
}

// 若 queue 为为空队列，则返回 TRUE，否则返回 FALSE
func (queue *QueueArray) IsEmpty() bool {
	return queue.length == 0
}
