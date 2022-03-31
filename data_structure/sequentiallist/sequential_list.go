package sequentiallist

type SequentialList struct {
	length   int
	capacity int
	data     *[]interface{}
}

func New(capacity int) *SequentialList {
	data := make([]interface{}, capacity)

	return &SequentialList{
		length:   0,
		capacity: capacity,
		data:     &data,
	}
}

// 判断线性表是否为空
func (list *SequentialList) IsEmpty() bool {
	return list.length == 0
}

// 判断线性表是否已满
func (list *SequentialList) IsFull() bool {
	return list.length == list.capacity
}

// 在下标 i 处插入元素 el，复杂度为 O(n)
func (list *SequentialList) Insert(index int, el interface{}) bool {
	if index < 0 || index > list.length || list.IsFull() {
		return false
	}
	for i := list.length - 1; i >= index; i-- {
		(*list.data)[i+1] = (*list.data)[i]
	}
	(*list.data)[index] = el
	list.length++
	return true
}

// 获取下标为 i 的元素
func (list *SequentialList) Get(index int) (interface{}, bool) {
	if index < 0 || index >= list.length {
		return nil, false
	}
	return (*list.data)[index], true
}

// 获取所有元素
func (list *SequentialList) GetAll() []interface{} {
	all := []interface{}{}
	for i := 0; i < list.length; i++ {
		all = append(all, (*list.data)[i])
	}
	return all
}

// 在下标 i 处删除元素，复杂度为 O(n)
func (list *SequentialList) Remove(index int) bool {
	if index < 0 || index >= list.length || list.IsEmpty() {
		return false
	}
	for i := index; i < list.length; i++ {
		(*list.data)[i] = (*list.data)[i+1]
	}
	list.length--
	return true
}
