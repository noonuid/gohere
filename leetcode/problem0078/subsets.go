package problem0078

// 使用二进制数中的 0/1 序列指示 nums 中的数是否在子集中
func subsets(nums []int) [][]int {
	var n = 1 << len(nums)
	var sets = make([][]int, n)

	for mask := 0; mask < n; mask++ {
		var set = []int{}

		for index, num := range nums {
			// 根据索引由低到高地提取 mask 中的二进制位
			if mask>>index&1 == 1 {
				set = append(set, num)
			}
		}

		sets[mask] = set
	}

	return sets
}
