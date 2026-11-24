package problem0089

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(n int) []int) {
	testCases := []struct {
		n    int
		want []int
	}{
		{
			n:    2,
			want: []int{0, 1, 3, 2},
		},
		{
			n:    1,
			want: []int{0, 1},
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

func TestGrayCode(t *testing.T) {
	test(t, grayCode)
}

func TestGrayCode_formula(t *testing.T) {
	test(t, grayCode_formula)
}
