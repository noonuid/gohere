package problem0392

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(s string, t string) bool) {
	testCases := []struct {
		s    string
		t    string
		want bool
	}{
		{
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.s, testCase.t)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestIsSubsequence(t *testing.T) {
	test(t, isSubsequence)
}

func TestIsSubsequence_two_pointer(t *testing.T) {
	test(t, isSubsequence_two_pointer)
}

func TestIsSubsequence_preprocess(t *testing.T) {
	test(t, isSubsequence_preprocess)
}
