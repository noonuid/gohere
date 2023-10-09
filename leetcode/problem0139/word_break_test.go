package problem0139

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s string, wordDict []string) bool) {
	// 测试用例。
	testCases := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.s, testCase.wordDict)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestWordBreak(t *testing.T) {
	testFramework(t, wordBreak)
}
