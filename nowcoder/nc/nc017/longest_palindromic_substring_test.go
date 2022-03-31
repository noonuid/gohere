package nc017

import (
	"reflect"
	"testing"
)

func TestNc017(t *testing.T) {
	testCases := []struct {
		A        string
		expected int
	}{
		{
			A:        "babad",
			expected: 3,
		},
		{
			A:        "cbbd",
			expected: 2,
		},
		{
			A:        "a",
			expected: 1,
		},
		{
			A:        "ac",
			expected: 1,
		},
	}

	t.Run("getLongestPalindrome", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := getLongestPalindrome(testCase.A)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
