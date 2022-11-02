package problem0102

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var testCases = []struct {
	input []int
	want  [][]int
}{
	{
		input: []int{},
		want:  [][]int{},
	},
	{
		input: []int{1},
		want:  [][]int{{1}},
	},
	{
		input: []int{3, 9, 20, structure.NULL, structure.NULL, 15, 7},
		want:  [][]int{{3}, {9, 20}, {15, 7}},
	},
}

func TestLevelOrder(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		got := levelOrder(structure.Ints2TreeNode(testCase.input))

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
