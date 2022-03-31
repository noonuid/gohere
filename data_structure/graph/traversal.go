package graph

// 图的广度优先遍历
func (graph *Graph) BreadthFirstSearch() []int {
	var route = make([]int, 0, len(graph.vertexes))
	if len(graph.vertexes) == 0 {
		return route
	}

	// 记录顶点是否已被访问
	var visited = make(map[int]bool, len(graph.vertexes))
	// 队列，用于存储将要访问的顶点
	var queue = make([]int, 0, len(graph.vertexes))

	// 对未访问过的顶点执行依次遍历,若是连通图,只会执行一次
	for i := 0; i < len(graph.vertexes); i++ {
		// 每一轮的遍历起点
		var start = graph.vertexes[i]
		if !visited[start] {
			queue = append(queue, start)
			visited[start] = true

			for len(queue) > 0 {
				// 当前中心元素
				var cur = queue[0]
				queue = queue[1:]
				route = append(route, cur)

				// 将与当前顶点存在边的并且未访问过的顶点入队
				for anothor := range graph.edges[cur] {
					if !visited[anothor] {
						queue = append(queue, anothor)
						visited[anothor] = true
					}
				}
			}
		}
	}

	return route
}

// 图的深度优先遍历
func (graph *Graph) DepthFirstSearch() []int {
	var route = make([]int, 0, len(graph.vertexes))
	if len(graph.vertexes) == 0 {
		return route
	}

	// 记录顶点是否已被访问
	var visited = make(map[int]bool, len(graph.vertexes))
	// 栈，用于存储访问节点的顺序
	var stack = make([]int, 0, len(graph.vertexes))

	// 对未访问过的顶点执行依次遍历,若是连通图,只会执行一次
	for i := 0; i < len(graph.vertexes); i++ {
		// 每一轮的遍历起点
		var start = graph.vertexes[i]
		if !visited[start] {
			stack = append(stack, start)
			visited[start] = true

			for len(stack) > 0 {
				var cur = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				route = append(route, cur)

				// 将与当前顶点存在边的并且未访问过的顶点入栈
				for another := range graph.edges[cur] {
					if !visited[another] {
						stack = append(stack, another)
						visited[another] = true
					}
				}
			}
		}
	}

	return route
}
