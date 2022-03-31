package problem0322

import (
	"reflect"
	"testing"
)

func TestProblem0322(t *testing.T) {
	testCases := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			coins:    []int{1},
			amount:   1,
			expected: 1,
		},
		{
			coins:    []int{1},
			amount:   2,
			expected: 2,
		},
	}

	t.Run("coinChange", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := coinChange(testCase.coins, testCase.amount)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
