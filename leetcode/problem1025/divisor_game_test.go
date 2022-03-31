package problem1025

import (
	"reflect"
	"testing"
)

func TestProblem1758(t *testing.T) {
	t.Run("minOperations", func(t *testing.T) {
		testCases := []struct {
			input    int
			expected bool
		}{
			{
				input:    1,
				expected: false,
			},
			{
				input:    2,
				expected: true,
			},
			{
				input:    3,
				expected: false,
			},
		}

		for _, testCase := range testCases {
			actual := divisorGame(testCase.input)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("actual: %v, expected: %v", actual, testCase.expected)
			}
		}
	})
}
