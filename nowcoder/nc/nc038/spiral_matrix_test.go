package nc038

import (
	"reflect"
	"testing"
)

func TestNc038(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected []int
	}{
		{
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	t.Run("spiralOrder", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := spiralOrder(testCase.matrix)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
