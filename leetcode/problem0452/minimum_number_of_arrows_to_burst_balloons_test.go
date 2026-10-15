package problem0452

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(points [][]int) int) {
	testCases := []struct {
		points [][]int
		want   int
	}{
		{
			points: [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			want:   2,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.points)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestFindMinArrowShots(t *testing.T) {
	test(t, findMinArrowShots)
}

func TestFindMinArrowShots_greedy(t *testing.T) {
	test(t, findMinArrowShots_greedy)
}
