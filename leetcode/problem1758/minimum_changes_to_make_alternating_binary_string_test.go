package problem1758

import (
	"reflect"
	"testing"
)

func TestProblem1758(t *testing.T) {
	t.Run("minOperations", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected int
		}{
			{
				input:    "0100",
				expected: 1,
			},
			{
				input:    "10",
				expected: 0,
			},
			{
				input:    "1111",
				expected: 2,
			},
		}

		for _, testCase := range testCases {
			actual := minOperations(testCase.input)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("actual: %d, expected: %d", actual, testCase.expected)
			}
		}
	})
}
