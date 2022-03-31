package problem0005

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	s    string
	want string
}{
	{
		s:    "babad",
		want: "bab",
	},
	{
		s:    "cbbd",
		want: "bb",
	},
	{
		s:    "a",
		want: "a",
	},
	{
		s:    "ac",
		want: "a",
	},
}

func TestLongestPalindrome(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var got = longestPalindrome(testCase.s)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %s, want: %s",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLongestPalindromeDynamicProgramming(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var got = longestPalindromeDynamicProgramming(testCase.s)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %s, want: %s",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLongestPalindromeCenter(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var got = longestPalindromeCenter(testCase.s)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %s, want: %s",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
