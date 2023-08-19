package problem0039

// 39. 组合总和

// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
// 找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。

// 回溯法。
func combinationSum(candidates []int, target int) [][]int {
	canLen := len(candidates)
	res := make([][]int, 0)
	path := make([]int, 0)

	var recursion func(index int, sum int)
	recursion = func(index int, sum int) {
		if index >= canLen || sum >= target {
			if sum == target {
				newOne := make([]int, len(path))
				copy(newOne, path)
				res = append(res, newOne)
			}
			return
		}
		// 不使用 index 处的数。
		recursion(index+1, sum)
		// 使用 index 处的数。
		path = append(path, candidates[index])
		recursion(index, sum+candidates[index])
		path = path[:len(path)-1]
	}

	recursion(0, 0)
	return res
}
