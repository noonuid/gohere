package problem0210

func findOrder_bfs(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0, numCourses)
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		head, tail := prerequisite[0], prerequisite[1]
		edges[tail] = append(edges[tail], head)
		indeg[head]++
	}

	queue := []int{}
	for vertex := range numCourses {
		if indeg[vertex] == 0 {
			queue = append(queue, vertex)
		}
	}
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)
		for _, head := range edges[vertex] {
			indeg[head]--
			if indeg[head] == 0 {
				queue = append(queue, head)
			}
		}
	}

	if len(result) < numCourses {
		return []int{}
	}
	return result
}

func findOrder_dfs(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0, numCourses)
	edges := make([][]int, numCourses)
	state := make([]int, numCourses)
	valid := true

	var dfs func(vertex int)
	dfs = func(vertex int) {
		state[vertex] = 1
		for _, head := range edges[vertex] {
			switch state[head] {
			case 0:
				dfs(head)
				if !valid {
					return
				}
			case 1:
				valid = false
				return
			}
		}
		state[vertex] = 2
		result = append(result, vertex)
	}

	for _, prerequisite := range prerequisites {
		head, tail := prerequisite[0], prerequisite[1]
		edges[tail] = append(edges[tail], head)
	}

	for vertex := range numCourses {
		if state[vertex] == 0 {
			dfs(vertex)
		}
		if !valid {
			return []int{}
		}
	}

	if !valid {
		return []int{}
	}

	for i, j := 0, numCourses-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
