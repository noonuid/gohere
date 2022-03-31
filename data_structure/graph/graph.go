package graph

type Graph struct {
	// 记录图的顶点
	vertexes []int
	// 记录图的边及其权（如果有的话，没有则值为 0），
	// map 的键为节点本身，而不是节点在 vertexes 中的索引
	edges map[int]map[int]int
	// 是否为有向图
	directed bool
}

// 初始化新的图
// directed 表示是否为有向图
func NewGraph(directed bool) *Graph {
	return &Graph{
		vertexes: []int{},
		edges:    make(map[int]map[int]int),
		directed: directed,
	}
}

// 在图中添加新的顶点
func (graph *Graph) AddVertex(vertex int) {
	if graph.edges == nil {
		graph.edges = make(map[int]map[int]int)
	}

	// 若顶点不存在则添加，存在则不做任何操作
	if _, ok := graph.edges[vertex]; !ok {
		graph.vertexes = append(graph.vertexes, vertex)

		graph.edges[vertex] = make(map[int]int)
	}
}

// 在图中批量添加新的顶点，该方法可确保按照指定顺序添加节点
func (graph *Graph) AddVertexes(vertexes []int) {
	for _, vertex := range vertexes {
		graph.AddVertex(vertex)
	}
}

// 在图中添加新的不带权重的边
func (graph *Graph) AddEdge(one, two int) {
	graph.AddWeightedEdge(one, two, 0)
}

// 在图中添加新的带权重的边
func (graph *Graph) AddWeightedEdge(one, two int, weight int) {
	// 若边的两个顶点没有添加，则先添加顶点
	graph.AddVertex(one)
	graph.AddVertex(two)

	graph.edges[one][two] = weight
	if !graph.directed {
		graph.edges[two][one] = weight
	}
}
