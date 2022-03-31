package graph

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	var vertexes = []int{1, 2, 3, 4, 5, 6}
	var edges = [][]int{
		{1, 2, 0},
		{1, 3, 0},
		{2, 4, 0},
		{3, 4, 0},
		{4, 5, 0},
		{5, 6, 0},
	}

	// Initializing graph, adding edges
	var graph = NewGraph(false)
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}

	var route = graph.BreadthFirstSearch()

	// 测试被遍历顶点的数量
	if len(vertexes) != len(route) {
		t.Errorf("Number of traverse, Expected: %d, Computed: %d", len(vertexes), len(route))
	}

	// 测试顶点是否都被遍历
	for i := 0; i < len(vertexes); i++ {
		var cur = vertexes[i]
		var exist = false

		for j := 0; j < len(route); j++ {
			if route[j] == cur {
				exist = true
				break
			}
		}

		if !exist {
			t.Errorf("顶点 %d 不存在", cur)
		}
	}

}

func TestDepthFirstSearch(t *testing.T) {
	var vertexes = []int{1, 2, 3, 4, 5, 6}
	var edges = [][]int{
		{1, 2, 0},
		{1, 3, 0},
		{2, 4, 0},
		{3, 4, 0},
		{4, 5, 0},
		{5, 6, 0},
	}

	// Initializing graph, adding edges
	var graph = NewGraph(false)
	for _, edge := range edges {
		graph.AddEdge(edge[0], edge[1])
	}

	var route = graph.DepthFirstSearch()

	// 测试被遍历顶点的数量
	if len(vertexes) != len(route) {
		t.Errorf("Number of traverse, Expected: %d, Computed: %d", len(vertexes), len(route))
	}

	// 测试顶点是否都被遍历
	for i := 0; i < len(vertexes); i++ {
		var cur = vertexes[i]
		var exist = false

		for j := 0; j < len(route); j++ {
			if route[j] == cur {
				exist = true
				break
			}
		}

		if !exist {
			t.Errorf("顶点 %d 不存在", cur)
		}
	}
}
