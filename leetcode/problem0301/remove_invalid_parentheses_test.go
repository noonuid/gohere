package problem0301

import (
	"reflect"
	"sort"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s string) []string) {
	// 测试用例。
	testCases := []struct {
		s    string
		want []string
	}{
		{
			s:    "()())()",
			want: []string{"(())()", "()()()"},
		},
		{
			s:    "(a)())()",
			want: []string{"(a())()", "(a)()()"},
		},
		{
			s:    ")(",
			want: []string{""},
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.s)
		sort.Strings(got)
		sort.Strings(testCase.want)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ntest case: %d\ngot : %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestRremoveInvalidParentheses_backtracking(t *testing.T) {
	testFramework(t, removeInvalidParentheses_backtracking)
}

func TestRremoveInvalidParentheses_bfs(t *testing.T) {
	testFramework(t, removeInvalidParentheses_bfs)
}
