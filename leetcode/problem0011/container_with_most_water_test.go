package problem0011

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) int) {
	testCases := []struct {
		height []int
		want   int
	}{
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			height: []int{1, 1},
			want:   1,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.height)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestMaxArea(t *testing.T) {
	testFramework(t, maxArea)
}
