package backtracking

// 计算一组数的全排列
func permute(nums []int) [][]int {
	var result = make([][]int, 0)
	backtracking(&result, []int{}, nums)

	return result
}

// 回溯法
// result 为最终的结果集
// path 为求解过程中的不完全排列结果
// unused 为排列过程中还未使用过的数
func backtracking(result *[][]int, path, unused []int) {
	// 所有的数均以使用，一次完整的排列完成
	if len(unused) == 0 {
		// 为 path 开辟新的空间，断开与之前空间的引用
		path = append([]int{}, path...)
		*result = append(*result, path)

		return
	}

	// 依次选择 nums 中的数放置在新的位置
	for i := 0; i < len(unused); i++ {
		// 选择 i 处的数作为新的位置的结果
		var cur = unused[i]
		path = append(path, cur)

		// i 处的数已经使用，从 nums 中移除
		unused = append(unused[:i], unused[i+1:]...)

		// 继续完成后序位置的排列
		backtracking(result, path, unused)

		// 返回本层之后恢复本层的 path, unused
		path = path[:len(path)-1]
		unused = append(unused[:i], append([]int{cur}, unused[i:]...)...)
	}
}
