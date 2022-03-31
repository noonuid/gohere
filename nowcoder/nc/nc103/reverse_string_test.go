package nc103

import (
	"reflect"
	"testing"
)

func TestNc103(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{
			str:      "abcd",
			expected: "dcba",
		},
		{
			str:      "",
			expected: "",
		},
	}

	t.Run("solve", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := solve(testCase.str)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
