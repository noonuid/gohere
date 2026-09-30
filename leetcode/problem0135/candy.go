package problem0135

func candy(ratings []int) int {
	ans := 0
	n := len(ratings)
	lefts := make([]int, n)
	lefts[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			lefts[i] = lefts[i-1] + 1
		} else {
			lefts[i] = 1
		}
	}
	right := 1
	ans += max(lefts[n-1], right)
	for i := n - 2; i > -1; i-- {
		if ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(lefts[i], right)
	}
	return ans
}

func candy_one_traversal(ratings []int) int {
	n := len(ratings)
	ans, pre, inc, dec := 1, 1, 1, 0
	for i := 1; i < n; i++ {
		if ratings[i-1] <= ratings[i] {
			if ratings[i-1] == ratings[i] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
			dec = 0
		} else {
			dec++
			if dec == inc {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}
