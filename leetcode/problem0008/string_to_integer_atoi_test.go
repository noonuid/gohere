package problem0008

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
			s:    "42",
			want: 42,
		},
		{
			s:    " -042",
			want: -42,
		},
		{
			s:    "1337c0d3",
			want: 1337,
		},
		{
			s:    "0-1",
			want: 0,
		},
		{
			s:    "words and 987",
			want: 0,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.s)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestMyAtoi(t *testing.T) {
	test(t, myAtoi)
}

func TestMyAtoi_automaton(t *testing.T) {
	test(t, myAtoi_automaton)
}
