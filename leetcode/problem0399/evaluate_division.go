package problem0399

// 399. 除法求值

// 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。

// 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。

// 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。

// 注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。

// 注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。

// 并查集。
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	equationsLen, queriesLen := len(equations), len(queries)

	parent, weight := make(map[string]string, 2*equationsLen), make(map[string]float64, 2*equationsLen)
	var find func(a string) string
	find = func(node string) string {
		if parent[node] != node {
			previousParent := parent[node]
			parent[node] = find(parent[node])
			weight[node] *= weight[previousParent]
		}
		return parent[node]
	}
	union := func(nodeA, nodeB string, value float64) string {
		parentA, parentB := find(nodeA), find(nodeB)
		if parentA == parentB {
			return parentA
		}
		parent[parentA] = parentB
		weight[parentA] = value * weight[nodeB] / weight[nodeA]
		return parentB
	}

	for i, equation := range equations {
		a, b := equation[0], equation[1]
		if _, ok := parent[a]; !ok {
			parent[a] = a
			weight[a] = 1.0
		}
		if _, ok := parent[b]; !ok {
			parent[b] = b
			weight[b] = 1.0
		}
		union(a, b, values[i])
	}

	answers := make([]float64, queriesLen)
	for i, query := range queries {
		a, b := query[0], query[1]
		_, okA := parent[a]
		_, okB := parent[b]
		if !okA || !okB {
			answers[i] = -1.0
		} else {
			parentA, parentB := find(a), find(b)
			if parentA != parentB {
				answers[i] = -1.0
			} else {
				answers[i] = weight[a] / weight[b]
			}
		}
	}
	return answers
}
