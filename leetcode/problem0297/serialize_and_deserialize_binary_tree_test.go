package problem0297

import (
	"reflect"
	"testing"

	"github.com/noonuid/go/leetcode/structure"
)

var null = structure.NULL

func TestProblem0297(t *testing.T) {
	// 测试用例。
	testCases := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: structure.Ints2Tree([]int{1, 2, 3, null, null, 4, 5}),
			want: []int{1, 2, 3, null, null, 4, 5},
		},
		{
			root: structure.Ints2Tree([]int{}),
			want: []int{},
		},
		{
			root: structure.Ints2Tree([]int{1}),
			want: []int{1},
		},
		{
			root: structure.Ints2Tree([]int{1, 2}),
			want: []int{1, 2},
		},
	}

	for caseIndex, testCase := range testCases {
		ser := Constructor()
		deser := Constructor()
		data := ser.serialize(testCase.root)
		ans := deser.deserialize(data)
		got := structure.Tree2Ints(ans)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}
