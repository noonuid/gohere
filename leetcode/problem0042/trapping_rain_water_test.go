package problem0042

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
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			height: []int{4, 2, 3},
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

func TestTrap(t *testing.T) {
	testFramework(t, trap)
}

func TestTrapTwoPointer(t *testing.T) {
	testFramework(t, trapTwoPointer)
}
