package problem0718

import (
	"reflect"
	"testing"
)

func TestProblem0718(t *testing.T) {
	testCases := []struct {
		nums1    []int
		nums2    []int
		expected int
	}{
		{
			nums1:    []int{0, 0, 0, 0, 0},
			nums2:    []int{0, 0, 0, 0, 0},
			expected: 5,
		},
		{
			nums1:    []int{1, 2, 3, 2, 1},
			nums2:    []int{3, 2, 1, 4, 7},
			expected: 3,
		},
		{
			nums1:    []int{0, 0, 0, 0, 1},
			nums2:    []int{1, 0, 0, 0, 0},
			expected: 4,
		},
		{
			nums1:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			nums2:    []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			expected: 59,
		},
	}

	t.Run("findLength", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := findLength(testCase.nums1, testCase.nums2)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("caseIndex: %d, case: %v, expected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
	t.Run("findLengthSlidingWindow", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := findLengthSlidingWindow(testCase.nums1, testCase.nums2)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("caseIndex: %d, case: %v, expected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
