package problem1859

import (
	"reflect"
	"testing"
)

func TestProblem1859(t *testing.T) {
	t.Run("sortSentence", func(t *testing.T) {
		testCases := []struct {
			input    string
			expected string
		}{
			{
				input:    "is2 sentence4 This1 a3",
				expected: "This is a sentence",
			},
			{
				input:    "Hi1",
				expected: "Hi",
			},
		}

		for _, testCase := range testCases {
			actual := sortSentence(testCase.input)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("actual: %s, expected: %s", actual, testCase.expected)
			}
		}
	})
}
