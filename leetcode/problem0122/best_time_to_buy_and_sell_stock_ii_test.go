package problem0122

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(prices []int) int) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.prices)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	testFramework(t, maxProfit)
}

func TestMaxProfit_greedy(t *testing.T) {
	testFramework(t, maxProfit_greedy)
}
