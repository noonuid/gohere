package problem1556

import (
	"reflect"
	"testing"
)

func TestProblem1556(t *testing.T) {
	testCases := []struct {
		n        int
		expected string
	}{
		{
			n:        987,
			expected: "987",
		},
		{
			n:        1234,
			expected: "1.234",
		},
		{
			n:        123456789,
			expected: "123.456.789",
		},
		{
			n:        0,
			expected: "0",
		},
	}

	t.Run("longestPalindrome", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := thousandSeparator(testCase.n)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %s, actual: %s",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
