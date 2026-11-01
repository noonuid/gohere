package problem0052

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
			n:    4,
			want: 2,
		},
		{
			n:    1,
			want: 1,
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

func TestTotalNQueens(t *testing.T) {
	test(t, totalNQueens)
}

func TestTotalNQueens_bitwise_operation(t *testing.T) {
	test(t, totalNQueens_bitwise_operation)
}
