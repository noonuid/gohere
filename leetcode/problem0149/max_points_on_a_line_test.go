package problem0149

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
			points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			want:   3,
		},
		{
			points: [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
			want:   4,
		},
		{
			points: [][]int{{0, 0}},
			want:   1,
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

func TestMaxPoints(t *testing.T) {
	test(t, maxPoints)
}
