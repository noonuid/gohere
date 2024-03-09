package problem0103

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var null = structure.NULL

func testFramework(t *testing.T, testFunc func(root *TreeNode) [][]int) {
	// 测试用例。
	testCases := []struct {
		root *TreeNode
		want [][]int
	}{
		{
			root: structure.Ints2Tree([]int{3, 9, 20, null, null, 15, 7}),
			want: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			root: structure.Ints2Tree([]int{1}),
			want: [][]int{{1}},
		},
		{
			root: structure.Ints2Tree([]int{}),
			want: [][]int{},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.root)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestZigzagLevelOrder(t *testing.T) {
	testFramework(t, zigzagLevelOrder)
}
