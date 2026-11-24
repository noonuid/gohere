package problem0089

func grayCode(n int) []int {
	ans := make([]int, 1, 1<<n)
	for i := 1; i <= n; i++ {
		num := 1 << (i - 1)
		for j := len(ans) - 1; j > -1; j-- {
			ans = append(ans, num|ans[j])
		}
	}
	return ans
}

func grayCode_formula(n int) []int {
	ans := make([]int, 1<<n)
	for i := 0; i < len(ans); i++ {
		ans[i] = i ^ (i >> 1)
	}
	return ans
}
