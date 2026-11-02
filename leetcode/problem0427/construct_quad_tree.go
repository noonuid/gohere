package problem0427

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)

	var merge func(top, left, bottom, right int) *Node
	merge = func(top, left, bottom, right int) *Node {
		if top == bottom && left == right {
			return &Node{
				Val:    grid[top][left] == 1,
				IsLeaf: true,
			}
		}

		midRow, midCol := (top+bottom)/2, (left+right)/2
		tl := merge(top, left, midRow, midCol)
		tr := merge(top, midCol+1, midRow, right)
		bl := merge(midRow+1, left, bottom, midCol)
		br := merge(midRow+1, midCol+1, bottom, right)

		node := &Node{}
		if (tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf) &&
			((tl.Val && tr.Val && bl.Val && br.Val) || !(tl.Val || tr.Val || bl.Val || br.Val)) {
			node.Val = tl.Val
			node.IsLeaf = true
		} else {
			node.TopLeft = tl
			node.TopRight = tr
			node.BottomLeft = bl
			node.BottomRight = br
		}
		return node
	}

	return merge(0, 0, n-1, n-1)
}

func construct_prefix_sum(grid [][]int) *Node {
	n := len(grid)
	preSum := make([][]int, n+1)
	preSum[0] = make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			preSum[i+1][j+1] = preSum[i+1][j] + preSum[i][j+1] - preSum[i][j] + grid[i][j]
		}
	}

	var dfs func(row0, col0, row1, col1 int) *Node
	dfs = func(row0, col0, row1, col1 int) *Node {
		sum := preSum[row1][col1] - preSum[row1][col0] - preSum[row0][col1] + preSum[row0][col0]
		if sum == 0 {
			return &Node{Val: false, IsLeaf: true}
		}
		if sum == (row1-row0)*(col1-col0) {
			return &Node{Val: true, IsLeaf: true}
		}

		rowMid, colMid := (row0+row1)/2, (col0+col1)/2
		return &Node{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     dfs(row0, col0, rowMid, colMid),
			TopRight:    dfs(row0, colMid, rowMid, col1),
			BottomLeft:  dfs(rowMid, col0, row1, colMid),
			BottomRight: dfs(rowMid, colMid, row1, col1),
		}
	}

	return dfs(0, 0, n, n)
}
