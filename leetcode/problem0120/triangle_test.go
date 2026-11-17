package problem0120

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(triangle [][]int) int) {
	testCases := []struct {
		triangle [][]int
		want     int
	}{
		{
			triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			want:     11,
		},
		{
			triangle: [][]int{{-10}},
			want:     -10,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.triangle)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestMinimumTotal(t *testing.T) {
	test(t, minimumTotal)
}
