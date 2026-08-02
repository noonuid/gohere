package problem0138

import (
	"fmt"
	"reflect"
	"testing"
)

// NULL 方便添加测试数据。
var NULL = -1 << 63

func list2Array(head *Node) [][]int {
	// 链条深度限制，链条深度超出此限制，会 panic。
	limit := 100

	times := 0

	result := [][]int{}
	for node := head; node != nil; node = node.Next {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过 %d，可能出现环状链条。请检查错误，或者放宽函数中 limit 的限制。", limit)
			panic(msg)
		}

		randomIndex := NULL
		if node.Random != nil {
			randomIndex = 0
			random := head
			for random != nil && random != node.Random {
				randomIndex++
				random = random.Next
			}
			if random == nil {
				randomIndex = NULL
			}
		}
		result = append(result, []int{node.Val, randomIndex})
	}

	return result
}

func array2List(array [][]int) *Node {
	var head *Node
	if len(array) == 0 {
		return head
	}
	head = &Node{Val: array[0][0]}
	node := head
	for i := 1; i < len(array); i++ {
		node.Next = &Node{Val: array[i][0]}
		node = node.Next
	}
	for i, node := 0, head; i < len(array); i, node = i+1, node.Next {
		if array[i][1] != NULL {
			randomIndex, random := 0, head
			for randomIndex < array[i][1] {
				randomIndex, random = randomIndex+1, random.Next
			}
			node.Random = random
		}
	}
	return head
}

func test(t *testing.T, fn func(head *Node) *Node) {
	testCases := []struct {
		input [][]int
		want  [][]int
	}{
		{
			input: [][]int{{7, NULL}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
			want:  [][]int{{7, NULL}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
		},
		{
			input: [][]int{{1, 1}, {2, 1}},
			want:  [][]int{{1, 1}, {2, 1}},
		},
		{
			input: [][]int{{3, NULL}, {3, 0}, {3, NULL}},
			want:  [][]int{{3, NULL}, {3, 0}, {3, NULL}},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(array2List(testCase.input))
		if !reflect.DeepEqual(list2Array(got), testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, list2Array(got), testCase.want)
		}
	}
}

func TestCopyRandomList(t *testing.T) {
	test(t, copyRandomList)
}

func TestCopyRandomList_hash_table(t *testing.T) {
	test(t, copyRandomList_hash_table)
}

func TestCopyRandomList_interleave_nodes(t *testing.T) {
	test(t, copyRandomList_interleave_nodes)
}
