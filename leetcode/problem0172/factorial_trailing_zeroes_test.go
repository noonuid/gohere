package problem0172

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(n int) int) {
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    3,
			want: 0,
		},
		{
			n:    5,
			want: 1,
		},
		{
			n:    0,
			want: 0,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.n)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestTrailingZeroes(t *testing.T) {
	test(t, trailingZeroes)
}

func TestTrailingZeroes_optimization(t *testing.T) {
	test(t, trailingZeroes_optimization)
}
