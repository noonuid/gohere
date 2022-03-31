package problem0209

import (
	"reflect"
	"testing"
)

func TestProblem0209(t *testing.T) {
	testCases := []struct {
		target   int
		nums     []int
		expected int
	}{
		{
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1,
		},
		{
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
	}

	t.Run("minSubArrayLen", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := minSubArrayLen(testCase.target, testCase.nums)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
