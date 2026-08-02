package problem0227

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(s string) int) {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "3+2*2",
			want: 7,
		},
		{
			s:    " 3/2 ",
			want: 1,
		},
		{
			s:    " 3+5 / 2 ",
			want: 5,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.s)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestCalculate(t *testing.T) {
	test(t, calculate)
}
