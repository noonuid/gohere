package nc126

import (
	"reflect"
	"testing"
)

func TestNc126(t *testing.T) {
	testCases := []struct {
		arr      []int
		aim      int
		expected int
	}{
		{
			arr:      []int{1, 2, 5},
			aim:      11,
			expected: 3,
		},
		{
			arr:      []int{2},
			aim:      3,
			expected: -1,
		},
		{
			arr:      []int{1},
			aim:      0,
			expected: 0,
		},
		{
			arr:      []int{1},
			aim:      1,
			expected: 1,
		},
		{
			arr:      []int{1},
			aim:      2,
			expected: 2,
		},
	}

	t.Run("minMoney", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := minMoney(testCase.arr, testCase.aim)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %v, actual: %v",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
