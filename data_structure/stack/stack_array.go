package stack

type StackArray struct {
	data []interface{}
}

// 初始化一个新栈
func NewStackArray() *StackArray {
	return &StackArray{data: []interface{}{}}
}

// 将元素入栈
func (stack *StackArray) Push(el interface{}) {
	stack.data = append(stack.data, el)
}

// 将元素出栈
func (stack *StackArray) Pop() interface{} {
	pop := stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return pop
}

// 返回栈的元素个数
func (stack *StackArray) Length() int {
	return len(stack.data)
}

// 判断栈是否为空
func (stack *StackArray) IsEmpty() bool {
	return len(stack.data) == 0
}
