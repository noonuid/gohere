package problem0188

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(k int, prices []int) int) {
	testCases := []struct {
		k      int
		prices []int
		want   int
	}{
		{
			k:      2,
			prices: []int{2, 4, 1},
			want:   2,
		},
		{
			k:      2,
			prices: []int{3, 2, 6, 5, 0, 3},
			want:   7,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.k, testCase.prices)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	test(t, maxProfit)
}
