package problem0151

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(s string) string) {
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			s:    "a good   example",
			want: "example good a",
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

func TestReverseWords(t *testing.T) {
	test(t, reverseWords)
}
