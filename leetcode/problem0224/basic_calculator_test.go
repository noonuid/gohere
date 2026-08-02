package problem0224

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
			s:    "1 + 1",
			want: 2,
		},
		{
			s:    " 2-1 + 2 ",
			want: 3,
		},
		{
			s:    "(1+(4+5+2)-3)+(6+8)",
			want: 23,
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
