package problem0123

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(prices []int) int) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			want:   6,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1},
			want:   0,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.prices)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	test(t, maxProfit)
}
