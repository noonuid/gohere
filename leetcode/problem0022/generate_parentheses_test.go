package problem0022

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(int) []string) {
	// 测试用例。
	testCases := []struct {
		n    int
		want []string
	}{
		{
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n:    1,
			want: []string{"()"},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.n)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestGenerateParenthesis_dfs(t *testing.T) {
	testFramework(t, generateParenthesis_dfs)
}

func TestGenerateParenthesis_brute(t *testing.T) {
	testFramework(t, generateParenthesis_brute)
}
