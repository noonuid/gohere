package problem0162

import (
	"slices"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int) int) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 1},
			want: []int{2},
		},
		{
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			want: []int{1, 5},
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.nums)
		if slices.Index(testCase.want, got) == -1 {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestFindPeakElement_traversal(t *testing.T) {
	testFramework(t, findPeakElement_max)
}

func TestFindPeakElement_binary(t *testing.T) {
	testFramework(t, findPeakElement_binary)
}
