package problem0240

import (
	"reflect"
	"testing"
)

func TestProblem0240(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			matrix:   [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target:   5,
			expected: true,
		},
		{
			matrix:   [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			target:   20,
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
	t.Run("searchMatrixBinarySearch", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := searchMatrixBinarySearch(testCase.matrix, testCase.target)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %t, actual: %t",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
	t.Run("searchMatrixZ", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := searchMatrixZ(testCase.matrix, testCase.target)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %t, actual: %t",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
