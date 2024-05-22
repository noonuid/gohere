package problem0013

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s string) int) {
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "III",
			want: 3,
		},
		{
			s:    "IV",
			want: 4,
		},
		{
			s:    "IX",
			want: 9,
		},
		{
			s:    "LVIII",
			want: 58,
		},
		{
			s:    "MCMXCIV",
			want: 1994,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.s)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestRomanToInt(t *testing.T) {
	testFramework(t, romanToInt)
}
