package problem0238

// 238. 除自身以外数组的乘积

// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

// 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。

// 动态规划，优化空间复杂度。
func productExceptSelf_dynamic_programming_space_optimization(nums []int) []int {
	n := len(nums)
	products := make([]int, n)
	products[0] = 1
	for i := 1; i < n; i++ {
		products[i] = products[i-1] * nums[i-1]
	}
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		products[i] = products[i] * rightProduct
		rightProduct = nums[i] * rightProduct
	}
	return products
}

// 动态规划。
func productExceptSelf_dynamic_programming(nums []int) []int {
	n := len(nums)
	products, leftProducts, rightProducts := make([]int, n), make([]int, n), make([]int, n)
	leftProducts[0], rightProducts[n-1] = 1, 1
	for i := 1; i < n; i++ {
		leftProducts[i] = leftProducts[i-1] * nums[i-1]
		rightProducts[n-1-i] = rightProducts[n-i] * nums[n-i]
	}
	for i := 0; i < n; i++ {
		products[i] = leftProducts[i] * rightProducts[i]
	}
	return products
}

// 暴力枚举。
// 超出时间限制。
func productExceptSelf_brute(nums []int) []int {
	n := len(nums)
	products := make([]int, n)
	for cur := 0; cur < n; cur++ {
		product := 1
		for i := 0; i < n; i++ {
			if i == cur {
				continue
			}
			product *= nums[i]
		}
		products[cur] = product
	}
	return products
}
