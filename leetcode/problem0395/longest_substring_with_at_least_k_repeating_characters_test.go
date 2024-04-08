package problem0395

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s string, k int) int) {
	// 测试用例。
	testCases := []struct {
		s    string
		k    int
		want int
	}{
		{
			s:    "aaabb",
			k:    3,
			want: 3,
		},
		{
			s:    "ababbc",
			k:    2,
			want: 5,
		},
		{
			s:    "aaaaaaaaabbbcccccddddd",
			k:    5,
			want: 10,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.s, testCase.k)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestLongestSubstring_enum(t *testing.T) {
	testFramework(t, longestSubstring_enum)
}

func TestLongestSubstring(t *testing.T) {
	testFramework(t, longestSubstring)
}
