package problem0224

func calculate(s string) int {
	n := len(s)
	result := 0
	sign := 1
	oprator := []int{1}
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = oprator[len(oprator)-1]
			i++
		case '-':
			sign = -oprator[len(oprator)-1]
			i++
		case '(':
			oprator = append(oprator, sign)
			i++
		case ')':
			oprator = oprator[:len(oprator)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			result = result + sign*num
		}
	}
	return result
}
