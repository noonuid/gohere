package problem0020

import (
	"reflect"
	"testing"
)

func TestProblem0020(t *testing.T) {
	testCases := []struct {
		s        string
		expected bool
	}{
		{
			s:        "()",
			expected: true,
		},
		{
			s:        "()[]{}",
			expected: true,
		},
		{
			s:        "(]",
			expected: false,
		},
		{
			s:        "([)]",
			expected: false,
		},
		{
			s:        "{[]}",
			expected: true,
		},
	}

	t.Run("isValid", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := isValid(testCase.s)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %t, actual: %t",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
