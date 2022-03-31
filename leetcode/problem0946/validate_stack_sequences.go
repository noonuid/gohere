package problem0946

func validateStackSequences(pushed []int, popped []int) bool {
	var stack = make([]int, 0)

	// 指示出栈元素的位置
	var index = 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])

		// 如果 popped[index] 与 pushed[i] 相等则一直出栈
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]

			index++
		}
	}

	return len(stack) == 0
}
