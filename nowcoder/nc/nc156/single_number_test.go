package nc156

import (
	"reflect"
	"testing"
)

func TestNc156(t *testing.T) {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{5, 4, 1, 1, 5, 1, 5},
			k:        3,
			expected: 4,
		},
		{
			nums:     []int{2, 2, 1},
			k:        2,
			expected: 1,
		},
	}

	t.Run("spiralOrder", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := foundOnceNumber(testCase.nums, testCase.k)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})

	t.Run("foundOnceNumberHashMap", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := foundOnceNumberHashMap(testCase.nums, testCase.k)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
