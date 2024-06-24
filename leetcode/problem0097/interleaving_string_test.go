package problem0097

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(s1 string, s2 string, s3 string) bool) {
	testCases := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
			want: true,
		},
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbbaccc",
			want: false,
		},
		{
			s1:   "",
			s2:   "",
			s3:   "",
			want: true,
		},
		{
			s1:   "",
			s2:   "",
			s3:   "a",
			want: false,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.s1, testCase.s2, testCase.s3)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestIsInterleave(t *testing.T) {
	test(t, isInterleave)
}
