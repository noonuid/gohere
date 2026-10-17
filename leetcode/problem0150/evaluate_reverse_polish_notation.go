package problem0150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			var value int
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				value = num1 + num2
			case "-":
				value = num1 - num2
			case "*":
				value = num1 * num2
			default:
				value = num1 / num2
			}
			stack = append(stack, value)
		}
	}
	return stack[0]
}

func evalRPN_array_stack(tokens []string) int {
	stack := make([]int, (len(tokens)+1)/2)
	index := -1
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			index++
			stack[index] = num
		} else {
			index--
			switch token {
			case "+":
				stack[index] += stack[index+1]
			case "-":
				stack[index] -= stack[index+1]
			case "*":
				stack[index] *= stack[index+1]
			default:
				stack[index] /= stack[index+1]
			}
		}
	}
	return stack[0]
}
