package problem0173

import (
	"reflect"
	"testing"

	"github.com/noonuid/gohere/leetcode/structure"
)

var null = structure.Null

func test(t *testing.T, fn func(root *TreeNode) BSTIterator) {
	testCases := []struct {
		funcs []string
		paras [][]int
		want  []any
	}{
		{
			funcs: []string{"BSTIterator", "next", "next", "hasNext", "next", "hasNext", "next", "hasNext", "next", "hasNext"},
			paras: [][]int{{7, 3, 15, null, null, 9, 20}, {}, {}, {}, {}, {}, {}, {}, {}, {}},
			want:  []any{nil, 3, 7, true, 9, true, 15, true, 20, false},
		},
	}

	for caseIndex, testCase := range testCases {
		var bstIterator BSTIterator
		got := []any{}
		for i, f := range testCase.funcs {
			switch f {
			case "BSTIterator":
				root := structure.Ints2Tree(testCase.paras[i])
				bstIterator = fn(root)
				got = append(got, nil)
			case "next":
				value := bstIterator.Next()
				got = append(got, value)
			case "hasNext":
				value := bstIterator.HasNext()
				got = append(got, value)
			}
		}

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestBSTIterator(t *testing.T) {
	test(t, Constructor)
}
