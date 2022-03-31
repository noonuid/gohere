package problem0344

import (
	"reflect"
	"testing"
)

func TestProblem0344(t *testing.T) {
	testCases := []struct {
		s        []byte
		expected []byte
	}{
		{
			s:        []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:        []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	t.Run("reverseString", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			reverseString(testCase.s)
			actual := testCase.s

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
