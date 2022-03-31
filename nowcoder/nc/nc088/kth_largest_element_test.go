package nc088

import (
	"reflect"
	"testing"
)

func TestNc088(t *testing.T) {
	// 测试用例，expected 字段为期望的返回值，其余字段为方法参数
	testCases := []struct {
		a        []int
		n        int
		K        int
		expected int
	}{
		{
			a:        []int{1, 3, 5, 2, 2},
			n:        5,
			K:        3,
			expected: 2,
		},
		{
			a:        []int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2},
			n:        12,
			K:        3,
			expected: 9,
		},
	}

	t.Run("findKth", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的实际返回值
			actual := findKth(testCase.a, testCase.n, testCase.K)

			// 如果方法的期望值与实际值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
