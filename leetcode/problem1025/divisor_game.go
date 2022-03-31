package problem1025

func divisorGame(n int) bool {
	f := make([]bool, n+5)

	f[1] = false
	f[2] = true
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !f[i-j] {
				f[i] = true
			}
		}
	}

	return f[n]
}
