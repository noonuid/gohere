package problem0215

import (
	"reflect"
	"testing"
)

func TestProblem0215(t *testing.T) {
	// 测试用例，expected 字段为期望的返回值，其余字段为方法参数
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
	}

	t.Run("findKthLargest", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的实际返回值
			actual := findKthLargest(testCase.nums, testCase.k)

			// 如果方法的期望值与实际值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})

	// 重置测试用例
	testCases = []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
	}
	t.Run("findKthLargestSelect", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的实际返回值
			actual := findKthLargestSelect(testCase.nums, testCase.k)

			// 如果方法的期望值与实际值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
