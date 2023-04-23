package problem0239

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int, int) []int) {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1},
		},
		{
			nums: []int{7, 2, 4},
			k:    2,
			want: []int{7, 4},
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.nums, testCase.k)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	testFramework(t, maxSlidingWindow)
}

func TestMaxSlidingWindow_brute_force(t *testing.T) {
	testFramework(t, maxSlidingWindow_brute_force)
}
