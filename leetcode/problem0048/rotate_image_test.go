package problem0048

import (
	"reflect"
	"testing"
)

func TestProblem0048(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			matrix:   [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			expected: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			matrix:   [][]int{{1}},
			expected: [][]int{{1}},
		},
		{
			matrix:   [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{3, 1}, {4, 2}},
		},
	}

	t.Run("rotate", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			rotate(testCase.matrix)

			if !reflect.DeepEqual(testCase.matrix, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, testCase.matrix)
			}
		}
	})

	// 重新初始化用例
	testCases = []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			matrix:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expected: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			matrix:   [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			expected: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			matrix:   [][]int{{1}},
			expected: [][]int{{1}},
		},
		{
			matrix:   [][]int{{1, 2}, {3, 4}},
			expected: [][]int{{3, 1}, {4, 2}},
		},
	}
	t.Run("rotateFlip", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			rotateFlip(testCase.matrix)

			if !reflect.DeepEqual(testCase.matrix, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, testCase.matrix)
			}
		}
	})
}
