package problem0136

import (
	"reflect"
	"testing"
)

func TestProblem0136(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
	}

	t.Run("spiralOrder", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := singleNumber(testCase.nums)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})

	t.Run("singleNumberXor", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := singleNumberXor(testCase.nums)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
