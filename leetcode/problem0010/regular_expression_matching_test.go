package problem0010

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(string, string) bool) {
	// 测试用例。
	testCases := []struct {
		s    string
		p    string
		want bool
	}{
		{
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			s:    "ab",
			p:    ".*",
			want: true,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.s, testCase.p)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestIsMatch(t *testing.T) {
	testFramework(t, isMatch)
}
