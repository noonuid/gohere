package graph

var Inf = 1<<31 - 1

// 图的 Dijkstra 最短路径算法
// start 为出发顶点
// dist 为出发顶点到其余各个顶点的最短路径长度
// path 为出发顶点到其余各个顶点的最短路径的对应前驱顶点
func (graph *Graph) Dijkstra(start int) (dist map[int]int, path map[int]int) {
	var length = len(graph.vertexes)
	if length == 0 {
		return
	}

	dist = make(map[int]int, length)
	path = make(map[int]int, length)
	// 记录顶点是否已经加入到最终的最短路径结果中
	var final = make(map[int]bool, length)
	// 初始化出发顶点到其余顶点的最短路径长度为 Inf
	for i := 0; i < length; i++ {
		dist[graph.vertexes[i]] = Inf
	}
	// 出发顶点到其自身的最短路径长度为 0
	dist[start] = 0

	for i := 0; i < length-1; i++ {
		// 获取当前已知的距离 start 最近的顶点
		var nearest = start
		var min = Inf
		for vertex, distance := range dist {
			if !final[vertex] && distance < min {
				min = distance
				nearest = vertex
			}
		}
		final[nearest] = true

		// 更新与 nearest 存在边的顶点的路径及距离
		for another, weight := range graph.edges[nearest] {
			if !final[another] && min+weight < dist[another] {
				dist[another] = min + weight
				path[another] = nearest
			}
		}
	}

	return
}

// 图的 Floyd 最短路径算法
// dist 为各个顶点的最短路径长度
// path 为各个顶点的最短路径的对应中转顶点
func (graph *Graph) Floyd() (dist map[int]map[int]int, path map[int]map[int]int) {
	var length = len(graph.vertexes)
	if length == 0 {
		return
	}

	dist = make(map[int]map[int]int, length)
	path = make(map[int]map[int]int, length)
	for i := 0; i < length; i++ {
		dist[graph.vertexes[i]] = make(map[int]int, length)
		path[graph.vertexes[i]] = make(map[int]int, length)

		for j := 0; j < length; j++ {
			if i == j {
				// 节点到自身的最短距离为 0
				dist[graph.vertexes[i]][graph.vertexes[j]] = 0
			} else if _, ok := graph.edges[graph.vertexes[i]][graph.vertexes[j]]; !ok {
				// 若两个节点之间没有边，则距离为无穷大
				dist[graph.vertexes[i]][graph.vertexes[j]] = Inf
			} else {
				// 存在边的两个节点之间的初始最短距离为边的权重
				dist[graph.vertexes[i]][graph.vertexes[j]] = graph.edges[graph.vertexes[i]][graph.vertexes[j]]
			}

			// 初始状态所有的两两节点之间都没有中转节点
			path[graph.vertexes[i]][graph.vertexes[j]] = -1
		}
	}

	for k := 0; k < length; k++ {
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if dist[graph.vertexes[i]][graph.vertexes[k]]+
					dist[graph.vertexes[k]][graph.vertexes[j]] <
					dist[graph.vertexes[i]][graph.vertexes[j]] {
					// 以 vertexes[k] 为中转节点更新 vertexes[i] vertexes[j] 之间的最短距离
					dist[graph.vertexes[i]][graph.vertexes[j]] =
						dist[graph.vertexes[i]][graph.vertexes[k]] +
							dist[graph.vertexes[k]][graph.vertexes[j]]

					path[graph.vertexes[i]][graph.vertexes[j]] = graph.vertexes[k]
				}
			}
		}
	}

	return
}
