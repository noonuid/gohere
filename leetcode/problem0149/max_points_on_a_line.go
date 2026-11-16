package problem0149

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	ans := 0
	for i := 0; i < n; i++ {
		if ans > n/2 || ans >= n-i {
			break
		}

		counts := make(map[int]int)
		for j := i + 1; j < n; j++ {
			dx, dy := points[i][0]-points[j][0], points[i][1]-points[j][1]
			if dx == 0 {
				dy = 1
			} else if dy == 0 {
				dx = 1
			} else {
				if dy < 0 {
					dx, dy = -dx, -dy
				}
				divisor := gcd(abs(dx), abs(dy))
				dx, dy = dx/divisor, dy/divisor
			}
			counts[dy+20001*dx]++
		}
		for _, count := range counts {
			if pCount := count + 1; pCount > ans {
				ans = pCount
			}
		}
	}
	return ans
}
