package problem0406

import "sort"

// 假设有打乱顺序的一群人站成一个队列，数组 people 表示队列中一些人的属性（不一定按顺序）。
// 每个 people[i] = [hi, ki] 表示第 i 个人的身高为 hi ，前面 正好 有 ki 个身高大于或等于 hi 的人。

// 请你重新构造并返回输入数组 people 所表示的队列。返回的队列应该格式化为数组 queue ，
// 其中 queue[j] = [hj, kj] 是队列中第 j 个人的属性（queue[0] 是排在队列前面的人）。

// 先排序，再贪心。
func reconstructQueue(people [][]int) [][]int {
	// 使 people 按 h 降序，k 升序排序。
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	for i, person := range people {
		// 将 person 插入到 index=k 处，k 为 person[1]。
		copy(people[person[1]+1:i+1], people[person[1]:i])
		people[person[1]] = person
	}
	return people
}

// 暴力枚举。超出时间限制。
func reconstructQueue_brute_force(people [][]int) [][]int {
	return backtrace([][]int{}, people)
}

func backtrace(path, unused [][]int) [][]int {
	if len(unused) == 0 {
		return path
	}
	isValid := func(path [][]int) bool {
		var count int
		for i := 0; i < len(path); i++ {
			count = 0
			for j := 0; j < i; j++ {
				if path[j][0] >= path[i][0] {
					count++
				}
			}
			if path[i][1] != count {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(unused); i++ {
		// 选择 unused[i]。
		path = append(path, unused[i])
		if isValid(path) {
			unusedNew := append(append([][]int{}, unused[:i]...), unused[i+1:]...)
			if result := backtrace(path, unusedNew); len(result) != 0 {
				return result
			}
		}
		// 还原。
		path = path[:len(path)-1]
	}

	return [][]int{}
}
