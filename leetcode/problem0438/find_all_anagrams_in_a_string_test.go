package problem0438

import (
	"reflect"
	"testing"
)

func TestProblem0438(t *testing.T) {
	t.Run("findAnagrams", func(t *testing.T) {
		testCases := []struct {
			s        string
			p        string
			expected []int
		}{
			{
				s:        "cbaebabacd",
				p:        "abc",
				expected: []int{0, 6},
			},
			{
				s:        "abab",
				p:        "ab",
				expected: []int{0, 1, 2},
			},
		}

		for _, testCase := range testCases {
			actual := findAnagrams(testCase.s, testCase.p)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("actual: %v, expected: %v", actual, testCase.expected)
			}
		}
	})
}
