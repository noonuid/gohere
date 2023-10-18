package problem0312

// 312. 戳气球

// 有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

// 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。
// 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

// 求所能获得硬币的最大数量。

// 动态规划。
func maxCoins_dynamic_programming(nums []int) int {
	n := len(nums) + 2
	// 向原有 nums 两端添加 1。
	reNums := make([]int, n)
	reNums[0], reNums[n-1] = 1, 1
	for i := 1; i < n-1; i++ {
		reNums[i] = nums[i-1]
	}
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	for left := n - 3; left > -1; left-- {
		for right := left + 2; right < n; right++ {
			for cur := left + 1; cur < right; cur++ {
				sum := reNums[left] * reNums[cur] * reNums[right]
				sum = sum + f[left][cur] + f[cur][right]
				if sum > f[left][right] {
					f[left][right] = sum
				}
			}
		}
	}
	return f[0][n-1]
}

// 每次向开区间 (left,right) 添加一个气球，并使用二维数组存储中间运算结果。
func maxCoins_memorized(nums []int) int {
	n := len(nums) + 2
	// 向原有 nums 两端添加 1。
	reNums := make([]int, n)
	reNums[0], reNums[n-1] = 1, 1
	for i := 1; i < n-1; i++ {
		reNums[i] = nums[i-1]
	}
	// mem 存储中间计算结果。
	mem := make([][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = -1
		}
	}
	var solve func(left, right int) int
	solve = func(left, right int) int {
		// 开区间 (left,right) 中没有元素，直接返回 0。
		if left >= right-1 {
			return 0
		}
		if mem[left][right] != -1 {
			return mem[left][right]
		}
		for cur := left + 1; cur < right; cur++ {
			sum := reNums[left] * reNums[cur] * reNums[right]
			sum = sum + solve(left, cur) + solve(cur, right)
			if sum > mem[left][right] {
				mem[left][right] = sum
			}
		}
		return mem[left][right]
	}
	return solve(0, n-1)
}

// 暴力枚举。
// 超出时间限制。
func maxCoins_brute(nums []int) int {
	max := len(nums)
	var recursion func(sum int, options []int)
	recursion = func(sum int, options []int) {
		n := len(options)
		if n == 0 {
			if sum > max {
				max = sum
			}
			return
		}
		for i := 0; i < n; i++ {
			left, cur, right := 1, options[i], 1
			if -1 < i-1 {
				left = options[i-1]
			}
			if i+1 < n {
				right = options[i+1]
			}
			options = append(options[:i], options[i+1:]...)
			recursion(sum+left*cur*right, options)
			options = append(options[:i], append([]int{cur}, options[i:]...)...)
		}
	}
	recursion(0, nums)
	return max
}
