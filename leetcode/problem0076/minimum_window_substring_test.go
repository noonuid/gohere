package problem0076

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s, t string) string) {
	// 测试用例
	testCases := []struct {
		s    string
		t    string
		want string
	}{
		{
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			s:    "a",
			t:    "aa",
			want: "",
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.s, testCase.t)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestMinWindow(t *testing.T) {
	testFramework(t, minWindow)
}

func TestMinWindow_brute_force(t *testing.T) {
	testFramework(t, minWindow_brute_force)
}
