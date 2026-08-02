package problem0227

func calculate(s string) int {
	n := len(s)
	sign := '+'
	stack := []int{}
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == n-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			num = 0
			sign = ch
		}
	}

	sum := 0
	for _, num := range stack {
		sum += num
	}
	return sum
}
