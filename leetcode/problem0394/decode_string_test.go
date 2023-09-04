package problem0394

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(string) string) {
	// 测试用例。
	testCases := []struct {
		s    string
		want string
	}{
		{
			s:    "3[a]2[bc]",
			want: "aaabcbc",
		},
		{
			s:    "3[a2[c]]",
			want: "accaccacc",
		},
		{
			s:    "2[abc]3[cd]ef",
			want: "abcabccdcdcdef",
		},
		{
			s:    "abc3[cd]xyz",
			want: "abccdcdcdxyz",
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.s)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestDecodeString_stack(t *testing.T) {
	testFramework(t, decodeString_stack)
}
