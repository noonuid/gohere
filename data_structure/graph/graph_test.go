package graph

import (
	"fmt"
	"testing"
)

var graphTestCases = []struct {
	// 测试用例的名称
	name string
	// 测试用例中图的边
	edges [][]int
	// 测试用例中图的顶点数量
	vertexNumber int
}{
	{
		name: "single edge",
		edges: [][]int{
			{0, 1, 1},
		},
		vertexNumber: 2,
	},
	{
		name: "many edges",
		edges: [][]int{
			{0, 1, 1},
			{0, 2, 2},
			{1, 3, 4},
			{3, 4, 3},
			{4, 8, 3},
			{4, 9, 1},
			{7, 8, 2},
			{8, 9, 2},
		},
		vertexNumber: 8,
	},
	{
		name: "cycles",
		edges: [][]int{
			{0, 1, 1},
			{0, 2, 2},
			{1, 3, 4},
			{3, 4, 3},
			{4, 2, 1},
		},
		vertexNumber: 5,
	},
	{
		name: "disconnected graphs",
		edges: [][]int{
			{0, 1, 5},
			{2, 4, 5},
			{3, 8, 5},
		},
		vertexNumber: 6,
	},
}

// 测试有向图
func TestDirectedGraph(t *testing.T) {
	// Testing self-loops separately only for directed graphs.
	// For undirected graphs each edge already creates a self-loop.
	var directedGraphTestCases = append(graphTestCases, struct {
		name         string
		edges        [][]int
		vertexNumber int
	}{
		name: "self-loops",
		edges: [][]int{
			{0, 1, 1},
			{1, 2, 2},
			{2, 1, 3},
		},
		vertexNumber: 3,
	})

	for _, test := range directedGraphTestCases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			// Initializing graph, adding edges
			var graph = NewGraph(true)
			for _, edge := range test.edges {
				graph.AddWeightedEdge(edge[0], edge[1], edge[2])
			}

			// 测试顶点的数量
			if len(graph.vertexes) != test.vertexNumber {
				t.Errorf("Number of vertexes, Expected: %d, Computed: %d", test.vertexNumber, len(graph.vertexes))
			}

			// 测试边的数量以及用例中的边是否存在
			var edgeCount = 0
			for _, e := range graph.edges {
				edgeCount += len(e)
			}
			if edgeCount != len(test.edges) {
				t.Errorf("Number of edges, Expected: %d, Computed: %d", len(test.edges), edgeCount)
			}
			for _, edge := range test.edges {
				if val, found := graph.edges[edge[0]][edge[1]]; !found || val != edge[2] {
					t.Errorf("Edge {%d->%d (%d)} not found", edge[0], edge[1], edge[2])
				}
			}
		})
	}
}

// 测试无向图
func TestUndirectedGraph(t *testing.T) {
	for _, test := range graphTestCases {
		t.Run(fmt.Sprint(test.name), func(t *testing.T) {
			// Initializing graph, adding edges
			var graph = NewGraph(false)
			for _, edge := range test.edges {
				graph.AddWeightedEdge(edge[0], edge[1], edge[2])
			}

			// 测试顶点的数量
			if len(graph.vertexes) != test.vertexNumber {
				t.Errorf("Number of vertexes, Expected: %d, Computed: %d", test.vertexNumber, len(graph.vertexes))
			}

			// 测试边的数量以及用例中的边是否存在
			var edgeCount = 0
			for _, e := range graph.edges {
				edgeCount += len(e)
			}
			if edgeCount != len(test.edges)*2 {
				t.Errorf("Number of edges, Expected: %d, Computed: %d", len(test.edges)*2, edgeCount)
			}
			for _, edge := range test.edges {
				if val, found := graph.edges[edge[0]][edge[1]]; !found || val != edge[2] {
					t.Errorf("Edge {%d->%d (%d)} not found", edge[0], edge[1], edge[2])
				}
			}
		})
	}
}
