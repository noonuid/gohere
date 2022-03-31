package problem1864

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
				input:    "100",
				expected: 1,
			},
			{
				input:    "111000",
				expected: 1,
			},
			{
				input:    "010",
				expected: 0,
			},
			{
				input:    "1110",
				expected: -1,
			},
		}

		for _, testCase := range testCases {
			actual := minSwaps(testCase.input)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("case: %v, actual: %d, expected: %d", testCase, actual, testCase.expected)
			}
		}
	})
}
