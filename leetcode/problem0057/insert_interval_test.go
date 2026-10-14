package problem0057

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(intervals [][]int, newInterval []int) [][]int) {
	testCases := []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.intervals, testCase.newInterval)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestInsert(t *testing.T) {
	test(t, insert)
}
