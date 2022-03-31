package nc055

import (
	"reflect"
	"testing"
)

func TestNc055(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected string
	}{
		{
			strs:     []string{"abca", "abc", "abca", "abc", "abcc"},
			expected: "abc",
		},
		{
			strs:     []string{"abc"},
			expected: "abc",
		},
	}

	t.Run("longestCommonPrefix", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := longestCommonPrefix(testCase.strs)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %s, actual: %s",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
