package problem0003

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	s    string
	want int
}{
	{
		s:    "abcabcbb",
		want: 3,
	},
	{
		s:    "bbbbb",
		want: 1,
	},
	{
		s:    "pwwkew",
		want: 3,
	},
	{
		s:    "",
		want: 0,
	},
	{
		s:    "dvdf",
		want: 3,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var got = lengthOfLongestSubstring(testCase.s)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLengthOfLongestSubstringMap(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var got = lengthOfLongestSubstringMap(testCase.s)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLengthOfLongestSubstringForce(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var got = lengthOfLongestSubstringForce(testCase.s)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
