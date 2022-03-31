package problem0020

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var pairs = map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack = []byte{}
	for i := 0; i < len(s); i++ {
		if right, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != right {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
