package problem1143

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(text1 string, text2 string) int) {
	// 测试用例。
	testCases := []struct {
		text1 string
		text2 string
		want  int
	}{
		{
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			text1: "abc",
			text2: "abc",
			want:  3,
		},
		{
			text1: "abc",
			text2: "def",
			want:  0,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.text1, testCase.text2)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	testFramework(t, longestCommonSubsequence)
}
