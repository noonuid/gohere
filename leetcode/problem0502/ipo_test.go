package problem0502

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(k int, w int, profits []int, capital []int) int) {
	testCases := []struct {
		k       int
		w       int
		profits []int
		capital []int
		want    int
	}{
		{
			k:       2,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{0, 1, 1},
			want:    4,
		},
		{
			k:       3,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{0, 1, 2},
			want:    6,
		},
		{
			k:       1,
			w:       0,
			profits: []int{1, 2, 3},
			capital: []int{1, 1, 2},
			want:    0,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.k, testCase.w, testCase.profits, testCase.capital)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestFindMaximizedCapital(t *testing.T) {
	test(t, findMaximizedCapital)
}
