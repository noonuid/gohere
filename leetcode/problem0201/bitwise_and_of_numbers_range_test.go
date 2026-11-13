package problem0201

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(left int, right int) int) {
	testCases := []struct {
		left  int
		right int
		want  int
	}{
		{
			left:  5,
			right: 7,
			want:  4,
		},
		{
			left:  0,
			right: 0,
			want:  0,
		},
		{
			left:  1,
			right: 2147483647,
			want:  0,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.left, testCase.right)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestRangeBitwiseAnd(t *testing.T) {
	test(t, rangeBitwiseAnd)
}

func TestRangeBitwiseAnd_clear_ones_after_common_prefix(t *testing.T) {
	test(t, rangeBitwiseAnd_clear_ones_after_common_prefix)
}
