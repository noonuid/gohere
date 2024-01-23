package problem0028

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(haystack string, needle string) int) {
	// 测试用例。
	testCases := []struct {
		haystack string
		needle   string
		want     int
	}{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
		{
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
	}
	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.haystack, testCase.needle)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ntest case: %d\ngot : %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestStrStr(t *testing.T) {
	testFramework(t, strStr)
}
