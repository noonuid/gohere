package problem0997

import (
	"reflect"
	"testing"
)

func TestProblem1859(t *testing.T) {
	t.Run("findJudge", func(t *testing.T) {
		testCases := []struct {
			n        int
			trust    [][]int
			expected int
		}{
			{
				n: 2,
				trust: [][]int{
					{1, 2},
				},
				expected: 2,
			},
			{
				n: 3,
				trust: [][]int{
					{1, 3},
					{2, 3},
				},
				expected: 3,
			},
			{
				n: 3,
				trust: [][]int{
					{1, 3},
					{2, 3},
					{3, 1},
				},
				expected: -1,
			},
			{
				n: 3,
				trust: [][]int{
					{1, 2},
					{2, 3},
				},
				expected: -1,
			},
			{
				n: 4,
				trust: [][]int{
					{1, 3},
					{1, 4},
					{2, 3},
					{2, 4},
					{4, 3},
				},
				expected: 3,
			},
		}

		for _, testCase := range testCases {
			actual := findJudge(testCase.n, testCase.trust)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("actual: %d, expected: %d", actual, testCase.expected)
			}
		}
	})
}
