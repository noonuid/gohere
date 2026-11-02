package problem0427

import (
	"testing"
)

// quadTreeToIntsSlice 将四叉树序列化为层序遍历的二维列表形式。
func quadTreeToIntsSlice(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*Node{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			res = append(res, nil)
		} else {
			isLeaf := 0
			if node.IsLeaf {
				isLeaf = 1
			}
			val := 0
			if node.Val {
				val = 1
			}
			res = append(res, []int{isLeaf, val})
			queue = append(queue, node.TopLeft, node.TopRight, node.BottomLeft, node.BottomRight)
		}
	}

	// 移除尾部的 nil，与标准测试用例的格式保持一致。
	for len(res) > 0 && res[len(res)-1] == nil {
		res = res[:len(res)-1]
	}

	return res
}

// equal 比较两个四叉树序列化结果是否相等。
func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] == nil && b[i] == nil {
			continue
		}
		if a[i] == nil || b[i] == nil {
			return false
		}
		if len(a[i]) != 2 || len(b[i]) != 2 {
			return false
		}
		// 比较 isLeaf 是否相等。
		if a[i][0] != b[i][0] {
			return false
		}
		// 如果 isLeaf 为 True (1)，则 val 必须相等。
		if a[i][0] == 1 {
			if a[i][1] != b[i][1] {
				return false
			}
		}
		// 当 isLeaf 为 False (0) 时，val 的值可以任意（True 或 False 均可被视为相等），因此不强制比较 a[i][1] 和 b[i][1]。
	}
	return true
}

func test(t *testing.T, fn func(grid [][]int) *Node) {
	testCases := []struct {
		grid [][]int
		want [][]int
	}{
		{
			grid: [][]int{
				{0, 1},
				{1, 0},
			},
			want: [][]int{
				{0, 1},
				{1, 0}, {1, 1}, {1, 1}, {1, 0},
			},
		},
		{
			grid: [][]int{
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
			},
			want: [][]int{
				{0, 1},
				{1, 1}, {0, 1}, {1, 1}, {1, 0},
				nil, nil, nil, nil, {1, 0}, {1, 0}, {1, 1}, {1, 1},
			},
		},
	}

	for caseIndex, testCase := range testCases {
		root := fn(testCase.grid)
		got := quadTreeToIntsSlice(root)

		if !equal(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestConstruct(t *testing.T) {
	test(t, construct)
}

func TestConstruct_prefix_sum(t *testing.T) {
	test(t, construct_prefix_sum)
}
