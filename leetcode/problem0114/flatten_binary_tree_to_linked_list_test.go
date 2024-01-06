package problem0114

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var null = structure.NULL

func testFramework(t *testing.T, testFunc func(root *TreeNode)) {
	// 测试用例。
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: structure.Ints2Tree([]int{1, 2, 5, 3, 4, null, 6}),
			want: []int{1, null, 2, null, 3, null, 4, null, 5, null, 6},
		},
		{
			root: structure.Ints2Tree([]int{}),
			want: []int{},
		},
		{
			root: structure.Ints2Tree([]int{0}),
			want: []int{0},
		},
	}

	for caseIndex, testCase := range testCases {
		testFunc(testCase.root)
		got := structure.Tree2Ints(testCase.root)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestFlatten_predecessor(t *testing.T) {
	testFramework(t, flatten_predecessor)
}

func TestFlatten_traverse_with_update(t *testing.T) {
	testFramework(t, flatten_traverse_with_update)
}

func TestFlatten_traverse_first(t *testing.T) {
	testFramework(t, flatten_traverse_first)
}
