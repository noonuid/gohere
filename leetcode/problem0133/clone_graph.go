package problem0133

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph_bfs(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	visited[node] = &Node{Val: node.Val, Neighbors: []*Node{}}
	queue := []*Node{node}
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		for _, neighbor := range head.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &Node{Val: neighbor.Val, Neighbors: []*Node{}}
				queue = append(queue, neighbor)
			}

			visited[head].Neighbors = append(visited[head].Neighbors, visited[neighbor])
		}
	}

	return visited[node]
}

func cloneGraph_dfs(node *Node) *Node {
	visited := make(map[*Node]*Node)

	var cg func(cur *Node) *Node
	cg = func(cur *Node) *Node {
		if cur == nil {
			return nil
		}

		var clonedNode *Node

		if clonedNode, ok := visited[cur]; ok {
			return clonedNode
		}

		clonedNode = &Node{Val: cur.Val, Neighbors: []*Node{}}
		visited[cur] = clonedNode

		for _, neighbor := range cur.Neighbors {
			clonedNode.Neighbors = append(clonedNode.Neighbors, cg(neighbor))
		}

		return clonedNode
	}

	return cg(node)
}
