package problem0074

import (
	"reflect"
	"testing"
)

func TestProblem0074(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   1,
			expected: true,
		},
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   60,
			expected: true,
		},
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
	}

	t.Run("searchMatrix", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := searchMatrix(testCase.matrix, testCase.target)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %t, actual: %t",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
	t.Run("searchMatrix1", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := searchMatrix1(testCase.matrix, testCase.target)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %t, actual: %t",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
