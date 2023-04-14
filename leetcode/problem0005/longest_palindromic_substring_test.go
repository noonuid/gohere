package problem0005

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(string) string) {
	// 测试用例
	testCases := []struct {
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

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.s)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLongestPalindrome(t *testing.T) {
	testFramework(t, longestPalindrome)
}

func TestLongestPalindromeDynamicProgramming(t *testing.T) {
	testFramework(t, longestPalindromeDynamicProgramming)
}

func TestLongestPalindromeCenter(t *testing.T) {
	testFramework(t, longestPalindromeCenter)
}
