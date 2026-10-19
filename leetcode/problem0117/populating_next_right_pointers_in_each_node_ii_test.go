package problem0117

import (
	"reflect"
	"testing"
)

var null = -1 << 63

// ints2Tree 根据 []int 生成树，并返回树的根节点。
// 借助队列实现。
func ints2Tree(ints []int) *Node {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &Node{
		Val: ints[0],
	}

	queue := make([]*Node, 1, n)
	queue[0] = root

	for i := 1; i < n && len(queue) > 0; {
		node := queue[0]
		queue = queue[1:]

		if ints[i] != null {
			node.Left = &Node{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != null {
			node.Right = &Node{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func test(t *testing.T, fn func(root *Node) *Node) {
	testCases := []struct {
		root *Node
		want []int
	}{
		{
			root: ints2Tree([]int{1, 2, 3, 4, 5, null, 7}),
			want: []int{1, null, 2, 3, null, 4, 5, 7, null},
		},
		{
			root: ints2Tree([]int{}),
			want: []int{},
		},
	}

	for caseIndex, testCase := range testCases {
		root := fn(testCase.root)

		head := root
		got := []int{}
		for head != nil {
			var nextHead *Node
			for node := head; node != nil; node = node.Next {
				got = append(got, node.Val)
				if cur := node.Left; cur != nil {
					if nextHead == nil {
						nextHead = cur
					}
				}
				if cur := node.Right; cur != nil {
					if nextHead == nil {
						nextHead = cur
					}
				}
			}
			got = append(got, null)
			head = nextHead
		}

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestConnect(t *testing.T) {
	test(t, connect)
}

func TestConnect_use_next(t *testing.T) {
	test(t, connect_use_next)
}
