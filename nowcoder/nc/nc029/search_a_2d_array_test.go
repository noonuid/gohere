package nc029

import (
	"reflect"
	"testing"
)

func TestNc029(t *testing.T) {
	testCases := []struct {
		target   int
		array    [][]int
		expected bool
	}{
		{
			array:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			array:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   1,
			expected: true,
		},
		{
			array:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   60,
			expected: true,
		},
		{
			array:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
	}

	t.Run("longestPalindrome", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := Find(testCase.target, testCase.array)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %t, actual: %t",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
