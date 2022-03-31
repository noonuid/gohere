package problem0012

import (
	"reflect"
	"testing"
)

func TestProblem0012(t *testing.T) {
	// 测试用例，expected 字段为期望的返回值，其余字段为方法参数
	testCases := []struct {
		num      int
		expected string
	}{
		{
			num:      3,
			expected: "III",
		},
		{
			num:      4,
			expected: "IV",
		},
		{
			num:      9,
			expected: "IX",
		},
		{
			num:      58,
			expected: "LVIII",
		},
		{
			num:      1994,
			expected: "MCMXCIV",
		},
	}

	t.Run("intToRoman", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的实际返回值
			actual := intToRoman(testCase.num)

			// 如果方法的期望值与实际值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
