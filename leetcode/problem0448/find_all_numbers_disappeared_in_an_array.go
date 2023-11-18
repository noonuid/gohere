package problem0448

// 448. 找到所有数组中消失的数字

// 给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

// 使用 nums 记录数字是否出现过。
func findDisappearedNumbers_nums(nums []int) []int {
	n := len(nums)
	for _, num := range nums {
		i := (num - 1) % n
		nums[i] += n
	}
	res := []int{}
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}
	return res
}

// 使用哈希表记录数字是否出现过。
func findDisappearedNumbers_map(nums []int) []int {
	n := len(nums)
	appeared := make(map[int]bool, n)
	for _, num := range nums {
		appeared[num] = true
	}
	res := []int{}
	for i := 1; i <= n; i++ {
		if !appeared[i] {
			res = append(res, i)
		}
	}
	return res
}
