package problem0207

// 207. 课程表

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

// 广度优先搜索。
func canFinish_bfs(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	indegrees := make([]int, numCourses)
	queue := []int{}
	count := 0
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indegrees[p[0]]++
	}
	for vertex, indegree := range indegrees {
		if indegree == 0 {
			queue = append(queue, vertex)
		}
	}
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		count++
		for _, adjacent := range edges[vertex] {
			indegrees[adjacent]--
			if indegrees[adjacent] == 0 {
				queue = append(queue, adjacent)
			}
		}
	}
	return count == numCourses
}

// 深度优先搜索。
func canFinish_dfs(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	valid := true
	var dfs func(vertex int)
	dfs = func(vertex int) {
		visited[vertex] = 1
		for _, adjacent := range edges[vertex] {
			if visited[adjacent] == 0 {
				dfs(adjacent)
				if !valid {
					return
				}
			} else if visited[adjacent] == 1 {
				valid = false
				return
			}
		}
		visited[vertex] = 2
	}
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}
	for vertex := 0; vertex < numCourses && valid; vertex++ {
		if visited[vertex] == 0 {
			dfs(vertex)
		}
	}
	return valid
}
