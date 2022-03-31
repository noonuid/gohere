package nc057

import (
	"reflect"
	"testing"
)

func TestNc057(t *testing.T) {
	testCases := []struct {
		x        int
		expected int
	}{
		{
			x:        123,
			expected: 321,
		},
		{
			x:        -123,
			expected: -321,
		},
		{
			x:        120,
			expected: 21,
		},
		{
			x:        0,
			expected: 0,
		},
	}

	t.Run("reverse", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := reverse(testCase.x)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
