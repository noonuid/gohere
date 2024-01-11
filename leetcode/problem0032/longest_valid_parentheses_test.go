package problem0032

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s string) int) {
	// 测试用例。
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "(()",
			want: 2,
		},
		{
			s:    ")()())",
			want: 4,
		},
		{
			s:    "",
			want: 0,
		},
		{
			s:    "()",
			want: 2,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.s)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestLongestValidParentheses_double_traversals(t *testing.T) {
	testFramework(t, longestValidParentheses_double_traversals)
}

func TestLongestValidParentheses_stack(t *testing.T) {
	testFramework(t, longestValidParentheses_stack)
}

func TestLongestValidParentheses_dynamic_programming(t *testing.T) {
	testFramework(t, longestValidParentheses_dynamic_programming)
}

func TestLongestValidParentheses_brute(t *testing.T) {
	testFramework(t, longestValidParentheses_brute)
}
