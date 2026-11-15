package problem0050

func myPow(x float64, n int) float64 {
	var recurse func(n int) float64
	recurse = func(n int) float64 {
		if n == 0 {
			return 1
		}

		y := recurse(n / 2)
		if n%2 == 0 {
			return y * y
		} else {
			return y * y * x
		}
	}

	if n >= 0 {
		return recurse(n)
	} else {
		return 1.0 / recurse(-n)
	}
}

func myPow_iterate(x float64, n int) float64 {
	iterate := func(n int) float64 {
		ans := 1.0
		contribution := x
		for n > 0 {
			if n%2 == 1 {
				ans *= contribution
			}
			contribution *= contribution
			n /= 2
		}
		return ans
	}

	if n >= 0 {
		return iterate(n)
	} else {
		return 1.0 / iterate(-n)
	}
}
