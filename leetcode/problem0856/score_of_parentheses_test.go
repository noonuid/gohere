package problem0856

import (
	"reflect"
	"testing"
)

func TestProblem0856(t *testing.T) {
	// 测试用例，expected 字段为期望的返回值，其余字段为方法参数
	testCases := []struct {
		s        string
		expected int
	}{
		{
			s:        "()",
			expected: 1,
		},
		{
			s:        "(())",
			expected: 2,
		},
		{
			s:        "()()",
			expected: 2,
		},
		{
			s:        "(()(()))",
			expected: 6,
		},
	}

	t.Run("scoreOfParentheses", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的实际返回值
			actual := scoreOfParentheses(testCase.s)

			// 如果方法的期望值与实际值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
	t.Run("scoreOfParenthesesStack", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的实际返回值
			actual := scoreOfParenthesesStack(testCase.s)

			// 如果方法的期望值与实际值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
