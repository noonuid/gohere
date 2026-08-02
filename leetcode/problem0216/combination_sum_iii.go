package problem0216

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	var backtrack func(path []int, sum int, start int)
	backtrack = func(path []int, sum int, start int) {
		if len(path) == k && sum == n {
			result = append(result, append([]int(nil), path...))
			return
		}

		for i := start; i < 10; i++ {
			backtrack(append(path, i), sum+i, i+1)
		}
	}
	backtrack([]int{}, 0, 1)
	return result
}

func combinationSum3_binary(k int, n int) [][]int {
	result := [][]int{}
	var combination []int
	sum := 0
	check := func(mask int) bool {
		combination = []int{}
		sum = 0
		for i := 0; i < 9; i++ {
			if (1<<i)&mask > 0 {
				combination = append(combination, i+1)
				sum += i + 1
			}
		}
		return len(combination) == k && sum == n
	}
	for mask := 0; mask < 1<<9; mask++ {
		if check(mask) {
			result = append(result, combination)
		}
	}
	return result
}
