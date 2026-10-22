package problem0133

import (
	"reflect"
	"sort"
	"testing"
)

func adjListToGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make(map[int]*Node, len(adjList))

	// 创建所有节点，确保邻居引用的是同一个节点对象。
	for i := range adjList {
		val := i + 1
		nodes[val] = &Node{Val: val}
	}

	// 建立节点之间的邻接关系。
	for i, neighbors := range adjList {
		node := nodes[i+1]
		for _, neighborVal := range neighbors {
			node.Neighbors = append(node.Neighbors, nodes[neighborVal])
		}
	}

	return nodes[1]
}

func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	visited := make(map[*Node]bool)
	queue := []*Node{node}
	visited[node] = true

	// 记录 BFS 遍历顺序。
	var nodes []*Node

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		nodes = append(nodes, cur)

		for _, neighbor := range cur.Neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	// 按节点值生成邻接表。
	adjList := make([][]int, len(nodes))
	for _, cur := range nodes {
		neighbors := make([]int, 0, len(cur.Neighbors))
		for _, neighbor := range cur.Neighbors {
			neighbors = append(neighbors, neighbor.Val)
		}
		// 按邻居节点的 Val 升序排序，确保无论遍历顺序如何，结果集顺序都保持一致。
		sort.Ints(neighbors)
		adjList[cur.Val-1] = neighbors
	}

	return adjList
}

func test(t *testing.T, fn func(node *Node) *Node) {
	testCases := []struct {
		adjList [][]int
		want    [][]int
	}{
		{
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			want:    [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			adjList: [][]int{},
			want:    [][]int{},
		},
	}

	for caseIndex, testCase := range testCases {
		node := adjListToGraph(testCase.adjList)
		clonedNode := fn(node)
		got := graphToAdjList(clonedNode)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestCloneGraph_bfs(t *testing.T) {
	test(t, cloneGraph_bfs)
}

func TestCloneGraph_dfs(t *testing.T) {
	test(t, cloneGraph_dfs)
}
