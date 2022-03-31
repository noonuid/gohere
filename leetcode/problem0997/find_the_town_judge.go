package problem0997

func findJudge(n int, trust [][]int) int {
	out := make([]int, n)
	in := make([]int, n)

	for _, pair := range trust {
		out[pair[0]-1]++
		in[pair[1]-1]++
	}

	for i := 0; i < n; i++ {
		if out[i] == 0 && in[i] == n-1 {
			return i + 1
		}
	}

	return -1
}
